package rest

import (
	"net/http"

	"github.com/ForeverSRC/todo-list-api/pkg/config"
	itemcreating "github.com/ForeverSRC/todo-list-api/pkg/item/creating"
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/item/listing"
	"github.com/gin-gonic/gin"
)

const (
	MsgSuccess = "success"

	CodeSuccess = 0
	CodeFail    = 500
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Handler(ic itemcreating.Service, il itemlisting.Service) *gin.Engine {
	gin.SetMode(config.Config.GetString("ginMode"))

	router := gin.Default()

	loadApiRouterGroup(router, ic, il)

	return router
}

func successJsonRes(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		successRes(data),
	)
}

func successRes(data interface{}) *Response {
	return &Response{
		Code:    CodeSuccess,
		Data:    data,
		Message: MsgSuccess,
	}
}

func errJsonRes(c *gin.Context, msg string) {
	c.JSON(
		http.StatusOK,
		errRes(msg),
	)
}

func errRes(msg string) *Response {
	return &Response{
		Code:    CodeFail,
		Data:    nil,
		Message: msg,
	}
}
