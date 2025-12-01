package handlers

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// HealthHandler 健康检查处理器
func HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":        "ok",
			"go_version":    runtime.Version(),
			"num_cpu":       runtime.NumCPU(),
			"num_goroutine": runtime.NumGoroutine(),
		})
	}
}
