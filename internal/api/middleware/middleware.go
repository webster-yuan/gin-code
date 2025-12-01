package middleware

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 记录接口耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "webster")
		// 1. 先记录要调用的函数名
		handlerName := c.HandlerName() // 例如 "main.getUser"
		path := c.Request.URL.Path     // 例如 "/user/index"

		c.Next() // 真正执行业务 handler
		cost := time.Since(start)

		fmt.Printf("[StatCost] %s | %s | %v\n", path, handlerName, cost)
	}
}

type BodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架ResponseWriter
	body               *bytes.Buffer // 记录用的response
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GinBodyLogMiddleware 记录响应体的中间件
func GinBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &BodyLogWriter{
			body:           bytes.NewBuffer([]byte{}),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyLogWriter
		c.Next()
		fmt.Printf("response body is : %v\n", bodyLogWriter.body.String())
	}
}
