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
	if message == "" {
		message = "操作成功！"
	}
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusOK, data, message})
}

func SendErrorJson(ctx *gin.Context, data any, errmessage string) {
	if errmessage == "" {
		errmessage = "操作失败！"
	}
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusBadRequest, data, errmessage})
}

func SendBadRequestJson(ctx *gin.Context, data any, errmessage string) {
	ctx.JSON(http.StatusBadRequest, ReturnMessage{http.StatusBadRequest, data, errmessage})
}

func SendNotFoundJson(ctx *gin.Context, data any, errmessage string) {
	if errmessage == "" {
		errmessage = "未找到相关数据！"
	}
	ctx.JSON(http.StatusOK, ReturnMessage{http.StatusNotFound, data, errmessage})
}

func SendJson(ctx *gin.Context, code int, data any, errmessage string) {
	ctx.JSON(code, ReturnMessage{code, data, errmessage})
}

func SendNotLoginJson(ctx *gin.Context, errmessage string) {
	if errmessage == "" {
		errmessage = "登录信息错误，请重新登陆！"
	}
	ctx.JSON(http.StatusUnauthorized, ReturnMessage{Code: http.StatusUnauthorized, Data: nil, Message: errmessage})
}
