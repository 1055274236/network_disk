package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// type responseWriter struct {
// 	gin.ResponseWriter
// 	b *bytes.Buffer // 记录返回的数据
// }

// // 重写Writer
// func (w responseWriter) Write(b []byte) (int, error) {
// 	w.b.Write(b)
// 	return w.ResponseWriter.Write(b)
// }

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// // 使用重写的writer
		// writer := responseWriter{
		// 	ctx.Writer,
		// 	bytes.NewBuffer([]byte{}),
		// }
		// ctx.Writer = writer

		timeNow := time.Now()
		ctx.Next()
		tc := time.Since(timeNow)
		fmt.Printf("%v\t%v\t%v\t%v\n", ctx.FullPath(), ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), tc)
	}
}
