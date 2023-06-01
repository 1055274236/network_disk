package middleware

import (
	"NetworkDisk/dao/oprationlogdao"
	"NetworkDisk/service"
	"NetworkDisk/utils/verifyuser"
	"bytes"
	"encoding/base64"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func UserVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			service.SendErrorJson(ctx, nil, "用户信息错误！请重新登陆！")
			ctx.Abort()
		}
		decodeString, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			service.SendErrorJson(ctx, nil, "用户信息错误！请重新登陆！")
			ctx.Abort()
		}
		user, err := verifyuser.DecodeUser(decodeString)
		if err != nil {
			service.SendErrorJson(ctx, nil, "用户信息错误！请重新登陆！")
			ctx.Abort()
		}

		timeNow := time.Now()
		b, _ := ctx.GetRawData()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		ctx.Next()
		go oprationlogdao.Add(user.Id, string(b), ctx.FullPath(), int(time.Since(timeNow).Milliseconds()), ctx.ClientIP())
	}
}
