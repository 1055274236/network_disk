package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SendSuccessJson(ctx *gin.Context, data any, message string) {
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusOK, data, message})
}

func SendErrorJson(ctx *gin.Context, data any, errmessage string) {
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusBadRequest, data, errmessage})
}

func SendJson(ctx *gin.Context, code int, data any, errmessage string) {
	ctx.JSON(code, ReturnMessage{code, data, errmessage})
}
