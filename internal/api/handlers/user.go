package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser 处理获取用户路由
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	}
}
