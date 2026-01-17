package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code      int         `json:"code"`                 // 业务状态码（与 HTTP 状态码一致）
	Message   string      `json:"message"`              // 消息描述
	Data      interface{} `json:"data,omitempty"`       // 数据（成功时返回）
	Error     string      `json:"error,omitempty"`      // 错误信息（失败时返回）
	Timestamp int64       `json:"timestamp"`            // 时间戳（Unix 时间戳，秒）
	RequestID string      `json:"request_id,omitempty"` // 请求 ID（用于追踪）
}

// Success 成功响应
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      http.StatusOK,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// SuccessWithCode 成功响应（自定义状态码）
func SuccessWithCode(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// Created 创建成功响应（201）
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Code:      http.StatusCreated,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	})
}

// NoContent 无内容响应（204）
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Error 错误响应
func Error(c *gin.Context, code int, message string, err error) {
	response := Response{
		Code:      code,
		Message:   message,
		Error:     message,
		Timestamp: time.Now().Unix(),
		RequestID: getRequestID(c),
	}

	// 如果有详细错误信息，在开发环境可以包含
	if err != nil && gin.Mode() == gin.DebugMode {
		response.Error = err.Error()
	}

	c.JSON(code, response)
}

// BadRequest 400 错误响应
func BadRequest(c *gin.Context, message string, err error) {
	Error(c, http.StatusBadRequest, message, err)
}

// Unauthorized 401 错误响应
func Unauthorized(c *gin.Context, message string, err error) {
	Error(c, http.StatusUnauthorized, message, err)
}

// Forbidden 403 错误响应
func Forbidden(c *gin.Context, message string, err error) {
	Error(c, http.StatusForbidden, message, err)
}

// NotFound 404 错误响应
func NotFound(c *gin.Context, message string, err error) {
	Error(c, http.StatusNotFound, message, err)
}

// InternalServerError 500 错误响应
func InternalServerError(c *gin.Context, message string, err error) {
	Error(c, http.StatusInternalServerError, message, err)
}

// getRequestID 获取请求 ID（可以从中间件中获取）
func getRequestID(c *gin.Context) string {
	requestID, exists := c.Get("request_id")
	if exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}
