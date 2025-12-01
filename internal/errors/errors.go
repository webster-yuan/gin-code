package errors

import (
	"errors"
	"fmt"
	"net/http"

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

		// 处理最后一个错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *AppError

			if errors.As(err, &appErr) {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
			} else {
				// 对于未处理的错误，返回500
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "内部服务器错误",
				})
			}
		}
	}
}
