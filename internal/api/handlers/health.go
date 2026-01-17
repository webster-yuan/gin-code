package handlers

import (
	"runtime"

	"gin/internal/api/response"

	"github.com/gin-gonic/gin"
)

// HealthHandler 健康检查处理器
func HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		healthData := gin.H{
			"status":        "ok",
			"go_version":    runtime.Version(),
			"num_cpu":       runtime.NumCPU(),
			"num_goroutine": runtime.NumGoroutine(),
		}
		response.Success(c, "服务运行正常", healthData)
	}
}
