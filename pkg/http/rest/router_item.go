package rest

import (
	"fmt"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/service/item/creating"
	itemdeleting "github.com/ForeverSRC/todo-list-api/pkg/service/item/deleting"
	itemediting "github.com/ForeverSRC/todo-list-api/pkg/service/item/editing"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/service/item/listing"
	itemmanaging "github.com/ForeverSRC/todo-list-api/pkg/service/item/managing"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
	"github.com/gin-gonic/gin"
)

func loadItemRouterGroup(router *gin.Engine, app *App) {
	groupA := router.Group("/api")
	{
		groupA.POST("/items", createItem(app.ItemCreator))
		groupA.GET("/items", listItems(app.ItemLister))
		groupA.PUT("/items/:id", editItem(app.ItemEditor))
		groupA.PUT("/items/:id/state", changeItemState(app.ItemManager))
		groupA.DELETE("/items/:id", deleteItem(app.ItemDeleter))
	}

}

func createItem(ic itemcreating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item model.ItemVo

		if err := c.ShouldBind(&item); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		id, err := ic.CreateItem(c.Request.Context(), item)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("create item error: %v", err))
			return
		}

		successJsonRes(c, gin.H{
			"itemId": id,
		})
	}
}

func listItems(il itemlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q vo.ItemListQuery
		if err := c.ShouldBindQuery(&q); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		list, err := il.ListItems(c.Request.Context(), &q)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("listing item error: %v", err))
			return
		}

		successJsonRes(c, list)
	}
}

func changeItemState(im itemmanaging.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			errJsonRes(c, "empty item id")
			return
		}

		var req vo.ItemManageRequest
		if err := c.ShouldBind(&req); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		if !req.State.ValidState() {
			errJsonRes(c, "invalid operation")
			return
		}

		req.Id = id

		if err := im.ChangeItemState(c.Request.Context(), &req); err != nil {
			errJsonRes(c, fmt.Sprintf("err: %v", err))
			return
		}

		successJsonRes(c, nil)
	}
}

func deleteItem(idel itemdeleting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			errJsonRes(c, "empty item id")
			return
		}

		err := idel.DeleteItem(c.Request.Context(), id)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("err: %v", err))
			return
		}

		successJsonRes(c, nil)

	}
}

func editItem(ie itemediting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			errJsonRes(c, "empty item id")
			return
		}

		var item model.ItemVo
		if err := c.ShouldBind(&item); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		err := ie.Edit(c.Request.Context(), id, item)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("err: %v", err))
			return
		}

		successJsonRes(c, nil)

	}
}
