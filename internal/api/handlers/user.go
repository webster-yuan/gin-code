package handlers

import (
	"fmt"
	"strconv"

	"gin/internal/api/response"
	"gin/internal/errors"
	"gin/internal/i18n"
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
// @Summary 创建新用户
// @Description 创建一个新的用户记录
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body models.CreateUserRequest true "用户信息"
// @Success 201 {object} response.Response{data=models.User} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未授权"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users [post]
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

		response.Created(c, i18n.UserMessage(i18n.UserCreateSuccess), user)
	}
}

// GetUser 获取用户
// @Summary 获取单个用户
// @Description 根据用户ID获取用户详细信息
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=models.User} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未授权"
// @Failure 404 {object} response.Response "用户不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(errors.NewBadRequestError(fmt.Sprintf(i18n.UserMessage(i18n.UserErrorInvalidID), idStr), err))
			return
		}

		user, err := h.userService.GetUserByID(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			return
		}

		response.Success(c, i18n.UserMessage(i18n.UserGetSuccess), user)
	}
}

// GetAllUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取系统中所有用户的列表
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]models.User} "获取成功"
// @Failure 401 {object} response.Response "未授权"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.userService.GetAllUsers(c.Request.Context())
		if err != nil {
			c.Error(err)
			return
		}

		response.Success(c, i18n.UserMessage(i18n.UserGetAllSuccess), users)
	}
}

// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 根据用户ID更新用户信息
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Param user body models.UpdateUserRequest true "更新的用户信息"
// @Success 200 {object} response.Response{data=models.User} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "用户不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(errors.NewBadRequestError(fmt.Sprintf(i18n.UserMessage(i18n.UserErrorInvalidID), idStr), err))
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

		response.Success(c, i18n.UserMessage(i18n.UserUpdateSuccess), user)
	}
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未授权"
// @Failure 404 {object} response.Response "用户不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Error(errors.NewBadRequestError(fmt.Sprintf(i18n.UserMessage(i18n.UserErrorInvalidID), idStr), err))
			return
		}

		err = h.userService.DeleteUser(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			return
		}

		response.Success(c, i18n.UserMessage(i18n.UserDeleteSuccess), nil)
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册新账户
// @Tags auth
// @Accept json
// @Produce json
// @Param register body models.CreateUserRequest true "注册信息"
// @Success 201 {object} response.Response{data=models.User} "注册成功"
// @Failure 400 {object} response.Response "请求参数错误或邮箱已被使用"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/auth/register [post]
func (h *UserHandler) Register() gin.HandlerFunc {
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

		response.Created(c, i18n.UserMessage(i18n.UserRegisterSuccess), user)
	}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并获取JWT令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "登录信息"
// @Success 200 {object} response.Response{data=models.LoginResponse} "登录成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "邮箱或密码错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		resp, err := h.userService.Login(c.Request.Context(), &req)
		if err != nil {
			c.Error(err)
			return
		}

		response.Success(c, i18n.UserMessage(i18n.UserLoginSuccess), resp)
	}
}

// RefreshToken 刷新访问令牌
// @Summary 刷新访问令牌
// @Description 使用刷新令牌获取新的访问令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body models.RefreshTokenRequest true "刷新令牌请求"
// @Success 200 {object} response.Response{data=models.RefreshTokenResponse} "刷新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "无效的刷新令牌"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/auth/refresh [post]
func (h *UserHandler) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RefreshTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		resp, err := h.userService.RefreshToken(c.Request.Context(), &req)
		if err != nil {
			c.Error(err)
			return
		}

		response.Success(c, i18n.UserMessage(i18n.UserRefreshTokenSuccess), resp)
	}
}
