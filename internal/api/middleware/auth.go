package middleware

import (
	"fmt"
	"strings"
	"time"

	"gin/internal/api/response"
	"gin/internal/auth"
	"gin/internal/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(jwtConfig *auth.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 添加日志，检查中间件是否被调用
		fmt.Println("=== 认证中间件被调用 ===")

		// 从请求头获取Authorization
		authHeader := c.GetHeader("Authorization")
		fmt.Printf("Authorization Header: %s\n", authHeader)

		if authHeader == "" {
			fmt.Println("未提供认证令牌，返回401")
			response.Unauthorized(c, "未提供认证令牌", nil)
			c.Abort()
			return
		}

		// 提取Bearer令牌
		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Printf("Token Parts: %v\n", parts)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			fmt.Println("认证令牌格式错误，返回401")
			response.Unauthorized(c, "认证令牌格式错误", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]
		fmt.Printf("Token String: %s\n", tokenString)

		// 验证令牌
		fmt.Println("开始验证令牌")
		claims, err := jwtConfig.ParseToken(tokenString)
		if err != nil {
			fmt.Printf("令牌验证失败: %v\n", err)
			response.Unauthorized(c, "无效的认证令牌", nil)
			c.Abort()
			return
		}

		fmt.Printf("令牌验证成功，用户ID: %d\n", claims.UserID)

		// 将用户信息存储在请求上下文中
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		fmt.Println("认证中间件执行完成")
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
