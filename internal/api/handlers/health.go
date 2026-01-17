package handlers

import (
	"runtime"

	"gin/internal/api/response"

	"github.com/gin-gonic/gin"
)

// HealthHandler 健康检查处理器
// @Summary 健康检查
// @Description 检查服务是否正常运行
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=map[string]interface{}} "服务正常"
// @Failure 500 {object} response.Response "服务异常"
// @Router /health [get]
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
