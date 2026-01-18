package middleware

import (
	"strings"
	"time"

	"gin/internal/api/response"
	"gin/internal/auth"
	"gin/internal/config"
	"gin/internal/i18n"
	"gin/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(jwtConfig *auth.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求ID（如果存在）
		requestID, _ := c.Get("request_id")
		requestIDStr := ""
		if id, ok := requestID.(string); ok {
			requestIDStr = id
		}

		path := c.Request.URL.Path
		method := c.Request.Method

		// 从请求头获取Authorization
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			logger.Log.Warn(i18n.LogMessage(i18n.LogAuthFailedNoToken),
				zap.String("request_id", requestIDStr),
				zap.String("path", path),
				zap.String("method", method),
			)
			response.Unauthorized(c, i18n.UserMessage(i18n.UserAuthNoToken), nil)
			c.Abort()
			return
		}

		// 提取Bearer令牌
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			logger.Log.Warn(i18n.LogMessage(i18n.LogAuthFailedInvalidFmt),
				zap.String("request_id", requestIDStr),
				zap.String("path", path),
				zap.String("method", method),
				zap.String("token_prefix", parts[0]),
			)
			response.Unauthorized(c, i18n.UserMessage(i18n.UserAuthInvalidFmt), nil)
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证令牌
		claims, err := jwtConfig.ParseToken(tokenString)
		if err != nil {
			logger.Log.Warn(i18n.LogMessage(i18n.LogAuthFailedInvalid),
				zap.String("request_id", requestIDStr),
				zap.String("path", path),
				zap.String("method", method),
				zap.Error(err),
			)
			response.Unauthorized(c, i18n.UserMessage(i18n.UserAuthInvalid), nil)
			c.Abort()
			return
		}

		// 将用户信息存储在请求上下文中
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		logger.Log.Debug(i18n.LogMessage(i18n.LogAuthSuccess),
			zap.String("request_id", requestIDStr),
			zap.String("path", path),
			zap.String("method", method),
			zap.Int64("user_id", claims.UserID),
			zap.String("email", claims.Email),
		)

		c.Next()
	}
}

// NewAuthMiddleware 创建认证中间件实例
func NewAuthMiddleware() gin.HandlerFunc {
	// 从配置获取JWT密钥和过期时间
	cfg := config.GetConfig()
	jwtConfig := auth.NewJWTConfig(
		cfg.JWT.SecretKey,
		time.Duration(cfg.JWT.ExpiresIn)*time.Hour,
	)

	return AuthMiddleware(jwtConfig)
}
