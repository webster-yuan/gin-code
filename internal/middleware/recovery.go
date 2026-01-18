package middleware

import (
	"gin/internal/i18n"
	"gin/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 自定义恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error(i18n.LogMessage(i18n.LogPanicRecovered),
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": i18n.UserMessage(i18n.UserErrorInternal),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
