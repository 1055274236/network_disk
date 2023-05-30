package middleware

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer // 记录返回的数据
}

// 重写Writer
func (w responseWriter) Write(b []byte) (int, error) {
	w.b.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用重写的writer
		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer

		timeNow := time.Now()
		c.Next()
		tc := time.Since(timeNow)
		fmt.Printf("%v\t%v\t%v\t%v\n", c.FullPath(), c.ClientIP(), time.Now().Format("2006-01-02"), tc)
	}
}
