package controller

import (
	"NetworkDisk/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"errmessage"`
}

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.Recovery(), middleware.Logger(), middleware.Cors())
}

func SendSuccessJson(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusOK, data, ""})
}

func SendErrorJson(ctx *gin.Context, data any, errmessage string) {
	ctx.JSON(http.StatusBadRequest, ReturnMessage{http.StatusBadRequest, data, errmessage})
}
