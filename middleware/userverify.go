package middleware

import (
	"NetworkDisk/config"
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
			service.SendNotLoginJson(ctx, "用户信息错误！请重新登陆！")
			ctx.Abort()
			return
		}
		decodeString, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			service.SendNotLoginJson(ctx, "用户信息错误！请重新登陆！")
			ctx.Abort()
			return
		}
		user, err := verifyuser.DecodeUser(decodeString)
		if err != nil {
			service.SendNotLoginJson(ctx, "用户信息错误！请重新登陆！")
			ctx.Abort()
			return
		}

		timeNow := time.Now()
		if user.CreatedAt+config.GlobalConfig.Gin.Login.Ext < timeNow.Unix() {
			service.SendNotLoginJson(ctx, "用户信息超时！请重新登陆！")
			ctx.Abort()
			return
		}

		ctx.Set("userId", user.Id)
		ctx.Set("account", user.Account)
		ctx.Next()
		ctx.Set("userId", nil)
		ctx.Set("account", nil)
		go oprationlogdao.Add(user.Id, "", ctx.FullPath(), int(time.Since(timeNow).Milliseconds()), ctx.ClientIP())
	}
}
