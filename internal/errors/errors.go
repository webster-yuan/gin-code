package errors

import (
	"errors"
	"fmt"
	"net/http"

	"gin/internal/api/response"

	"github.com/gin-gonic/gin"
)

// AppError 应用错误结构体
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"` // 不对外暴露原始错误
}

// Error 实现error接口
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// Unwrap 实现errors.Unwrap接口
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewBadRequestError 创建400错误
func NewBadRequestError(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}

// NewNotFoundError 创建404错误
func NewNotFoundError(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
		Err:     err,
	}
}

// NewInternalServerError 创建500错误
func NewInternalServerError(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Err:     err,
	}
}

// ErrorHandler 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果已经写入响应，不再处理
		if c.Writer.Written() {
			return
		}

		// 处理最后一个错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *AppError

			if errors.As(err, &appErr) {
				// 使用统一响应格式
				respondError(c, appErr)
			} else {
				// 对于未处理的错误，返回500
				respondError(c, NewInternalServerError("内部服务器错误", err))
			}
		}
	}
}

// respondError 使用统一响应格式返回错误
func respondError(c *gin.Context, appErr *AppError) {
	response.Error(c, appErr.Code, appErr.Message, appErr.Err)
}
