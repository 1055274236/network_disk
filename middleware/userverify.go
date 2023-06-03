package middleware

import (
	"NetworkDisk/dao/oprationlogdao"
	"NetworkDisk/service"
	"NetworkDisk/utils/verifyuser"
	"encoding/base64"
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
		b, ok := ctx.Get("ContextParams")
		if !ok {
			b = "err"
		}
		ctx.Set("account", user.Account)
		ctx.Next()
		ctx.Set("account", nil)
		go oprationlogdao.Add(user.Id, b.(string), ctx.FullPath(), int(time.Since(timeNow).Milliseconds()), ctx.ClientIP())
	}
}
