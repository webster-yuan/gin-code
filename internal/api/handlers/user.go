package handlers

import (
	"net/http"
	"strconv"

	"gin/internal/models"
	"gin/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		user, err := h.userService.CreateUser(c.Request.Context(), &req)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"code":    200,
			"message": "创建成功",
			"data":    user,
		})
	}
}

// GetUser 获取用户
func (h *UserHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		user, err := h.userService.GetUserByID(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    user,
		})
	}
}

// GetAllUsers 获取所有用户
func (h *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.userService.GetAllUsers(c.Request.Context())
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    users,
		})
	}
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		var req models.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		user, err := h.userService.UpdateUser(c.Request.Context(), id, &req)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新成功",
			"data":    user,
		})
	}
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		err = h.userService.DeleteUser(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}
}
