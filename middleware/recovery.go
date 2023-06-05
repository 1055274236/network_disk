package middleware

import (
	"NetworkDisk/dao/errlogdao"
	"NetworkDisk/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				contextParams, ok := ctx.Get("ContextParams")
				if !ok {
					contextParams = ""
				}

				hearder := ""
				for key, value := range ctx.Request.Header {
					hearder += key + ":" + strings.Join(value, ",") + "\n"
				}

				go errlogdao.Add(ctx.FullPath(), hearder, contextParams.(string), err.(string))

				service.SendBadRequestJson(ctx, err, err.(string))
			}
		}()
		ctx.Next()
	}
}
