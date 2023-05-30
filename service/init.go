package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"errmessage"`
}

func SendSuccessJson(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusOK, data, ""})
}

func SendErrorJson(ctx *gin.Context, data any, errmessage string) {
	ctx.JSON(http.StatusBadRequest, ReturnMessage{http.StatusBadRequest, data, errmessage})
}

func SendJson(ctx *gin.Context, code int, data any, errmessage string) {
	ctx.JSON(http.StatusOK, ReturnMessage{code, data, errmessage})
}
