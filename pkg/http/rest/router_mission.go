package rest

import (
	"fmt"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/service/item/creating"
	itemediting "github.com/ForeverSRC/todo-list-api/pkg/service/item/editing"
	missioncreating "github.com/ForeverSRC/todo-list-api/pkg/service/mission/creating"
	missiondetail "github.com/ForeverSRC/todo-list-api/pkg/service/mission/detail"
	missionitemadd "github.com/ForeverSRC/todo-list-api/pkg/service/mission/itemadd"
	missionlisting "github.com/ForeverSRC/todo-list-api/pkg/service/mission/listing"
	missionupdate "github.com/ForeverSRC/todo-list-api/pkg/service/mission/update"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
	"github.com/gin-gonic/gin"
)

func loadMissionRouterGroup(router *gin.Engine, app *App) {
	groupA := router.Group("/api")
	{
		groupA.POST("/missions", createMission(app.MissionCreator))
		groupA.GET("/missions", listMission(app.MissionLister))
		groupA.GET("/missions/:id", getMission(app.MissionDetail))
		groupA.PUT("/missions/:id", updateMission(app.MissionUpdater))
		groupA.POST("/missions/:id/items", addItemToMission(app.ItemCreator, app.ItemEditor, app.MissionItemAdder))
	}

}

func createMission(mc missioncreating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mv model.MissionVo
		if err := c.ShouldBind(&mv); err != nil {
			errJsonRes(c, fmt.Sprintf("create mission error: %v", err))
			return
		}

		id, err := mc.CreateMission(c.Request.Context(), mv)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("create mission error: %v", err))
			return
		}

		successJsonRes(c, gin.H{
			"missionId": id,
		})
	}
}

func listMission(ml missionlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q vo.MissionListQuery
		if err := c.ShouldBindQuery(&q); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		list, err := ml.ListMission(c.Request.Context(), q)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("listing mission error: %v", err))
			return
		}

		successJsonRes(c, list)
	}
}

func getMission(mg missiondetail.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			errJsonRes(c, "empty mission id")
			return
		}

		mission, err := mg.GetMission(c.Request.Context(), id)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("get mission [%s] error: %v", id, err))
			return
		}

		successJsonRes(c, gin.H{
			"mission": mission,
		})
	}
}

func updateMission(mu missionupdate.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			errJsonRes(c, "empty mission id")
			return
		}

		var value model.MissionVo
		if err := c.ShouldBind(&value); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		err := mu.UpdateMission(c.Request.Context(), id, value)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("update mission [%s] error: %v", id, err))
			return
		}

		successJsonRes(c, nil)

	}
}

func addItemToMission(ic itemcreating.Service, ie itemediting.Service, ma missionitemadd.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		mid := c.Param("id")
		if mid == "" {
			errJsonRes(c, "empty mission id")
			return
		}

		var req vo.MissionAddItemReq
		if err := c.ShouldBind(&req); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		switch req.AddType {
		case vo.AddNewItemType:
			req.ItemInfo.RelatedMission = mid
			itemId, err := ic.CreateItem(c.Request.Context(), *req.ItemInfo)
			if err != nil {
				errJsonRes(c, fmt.Sprintf("mission [%s] add item error: %v", mid, err))
				return
			}

			err = ma.AddItem(c.Request.Context(), mid, itemId)
			if err != nil {
				errJsonRes(c, fmt.Sprintf("mission [%s] add item[%s] error: %v", mid, itemId, err))
				return
			}
		case vo.AddExistItemType:
			item := model.ItemVo{
				RelatedMission: mid,
			}

			err := ie.Edit(c.Request.Context(), req.ItemId, item)
			if err != nil {
				errJsonRes(c, fmt.Sprintf("mission [%s] add item[%s] updating item error: %v", mid, req.ItemId, err))
			}

			err = ma.AddItem(c.Request.Context(), mid, req.ItemId)
			if err != nil {
				errJsonRes(c, fmt.Sprintf("mission [%s] add item[%s] updating mission error: %v", mid, req.ItemId, err))
				return
			}
		default:
			errJsonRes(c, "request param error")
			return
		}

		successJsonRes(c, nil)
	}
}
