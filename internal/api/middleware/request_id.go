package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware 请求ID中间件
// 为每个请求生成唯一的请求ID，用于日志追踪和问题排查
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从请求头获取请求ID（如果客户端已提供）
		requestID := c.GetHeader("X-Request-ID")

		// 如果没有，生成新的请求ID
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将请求ID存储到上下文中
		c.Set("request_id", requestID)

		// 将请求ID添加到响应头
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
