package rest

import (
	"fmt"

	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/item/creating"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/item/listing"
	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/gin-gonic/gin"
)

func loadApiRouterGroup(router *gin.Engine, ic itemcreating.Service, il itemlisting.Service) {
	groupA := router.Group("/api")
	{
		groupA.POST("/items", createItem(ic))
		groupA.GET("/items", listItems(il))
	}

}

func createItem(ic itemcreating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item model.Item

		if err := c.ShouldBind(&item); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		if err := ic.CreateItem(item); err != nil {
			errJsonRes(c, fmt.Sprintf("create item error: %v", err))
			return
		}

		successJsonRes(c, nil)
	}
}

func listItems(il itemlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q itemlisting.ItemListQuery
		if err := c.ShouldBindQuery(&q); err != nil {
			errJsonRes(c, fmt.Sprintf("binding error: %v", err))
			return
		}

		list, err := il.ListItems(&q)
		if err != nil {
			errJsonRes(c, fmt.Sprintf("listing item error: %v", err))
			return
		}

		successJsonRes(c, gin.H{
			"items": list,
		})
	}
}
