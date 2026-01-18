package middleware

import (
	"bytes"
	"time"

	"gin/internal/i18n"
	"gin/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

		// 获取请求ID（如果存在）
		requestID, _ := c.Get("request_id")
		requestIDStr := ""
		if id, ok := requestID.(string); ok {
			requestIDStr = id
		}

		logger.Log.Info(i18n.LogMessage(i18n.LogRequestCost),
			zap.String("request_id", requestIDStr),
			zap.String("path", path),
			zap.String("handler", handlerName),
			zap.Duration("cost", cost),
			zap.String("method", c.Request.Method),
		)
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

		// 获取请求ID（如果存在）
		requestID, _ := c.Get("request_id")
		requestIDStr := ""
		if id, ok := requestID.(string); ok {
			requestIDStr = id
		}

		// 记录响应体（仅在调试模式下或错误响应时记录）
		bodyStr := bodyLogWriter.body.String()
		if c.Writer.Status() >= 400 || logger.Log.Core().Enabled(zap.DebugLevel) {
			logger.Log.Debug(i18n.LogMessage(i18n.LogResponseBody),
				zap.String("request_id", requestIDStr),
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method),
				zap.Int("status", c.Writer.Status()),
				zap.String("body", bodyStr),
			)
		}
	}
}
