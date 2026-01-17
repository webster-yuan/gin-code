package service

import (
	"context"
	"fmt"

	"gin/internal/auth"
	"gin/internal/config"
	"gin/internal/errors"
	"gin/internal/models"
	"gin/internal/repository"
	"time"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, id int64, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	// 业务逻辑：检查邮箱是否已存在
	_, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil {
		// 邮箱已存在
		return nil, errors.NewBadRequestError("邮箱已被使用", fmt.Errorf("email already exists: %s", req.Email))
	}

	// 加密密码
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, errors.NewInternalServerError("密码加密失败", err)
	}

	// 创建用户
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      req.Age,
	}

	return s.userRepo.Create(ctx, user)
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	if id <= 0 {
		return nil, errors.NewBadRequestError("用户ID无效", fmt.Errorf("invalid user id: %d", id))
	}

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("用户不存在", err)
	}

	return user, nil
}

// GetUserByEmail 根据邮箱获取用户
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.NewBadRequestError("邮箱不能为空", fmt.Errorf("email is required"))
	}

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.NewNotFoundError("用户不存在", err)
	}

	return user, nil
}

// GetAllUsers 获取所有用户
func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("获取用户列表失败", err)
	}

	return users, nil
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(ctx context.Context, id int64, req *models.UpdateUserRequest) (*models.User, error) {
	if id <= 0 {
		return nil, errors.NewBadRequestError("用户ID无效", fmt.Errorf("invalid user id: %d", id))
	}

	// 先获取现有用户
	existingUser, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("用户不存在", err)
	}

	// 更新字段（只更新提供的字段）
	user := &models.User{
		ID:        existingUser.ID,
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		CreatedAt: existingUser.CreatedAt,
	}

	// 如果字段为空，使用原有值
	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}
	if user.Age == 0 {
		user.Age = existingUser.Age
	}

	// 如果邮箱有变化，检查新邮箱是否已被使用
	if user.Email != existingUser.Email {
		_, err := s.userRepo.FindByEmail(ctx, user.Email)
		if err == nil {
			return nil, errors.NewBadRequestError("邮箱已被使用", fmt.Errorf("email already exists: %s", user.Email))
		}
	}

	return s.userRepo.Update(ctx, id, user)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.NewBadRequestError("用户ID无效", fmt.Errorf("invalid user id: %d", id))
	}

	// 检查用户是否存在
	_, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return errors.NewNotFoundError("用户不存在", err)
	}

	return s.userRepo.Delete(ctx, id)
}

// Login 用户登录
func (s *userService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	// 检查邮箱是否存在
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NewUnauthorizedError("邮箱或密码错误", fmt.Errorf("invalid email or password"))
	}

	// 验证密码
	if !auth.CheckPassword(user.Password, req.Password) {
		return nil, errors.NewUnauthorizedError("邮箱或密码错误", fmt.Errorf("invalid email or password"))
	}

	// 获取JWT配置
	cfg := config.GetConfig()
	jwtConfig := auth.NewJWTConfig(
		cfg.JWT.SecretKey,
		time.Duration(cfg.JWT.ExpiresIn)*time.Hour,
	)

	// 生成JWT令牌
	token, err := jwtConfig.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, errors.NewInternalServerError("生成令牌失败", err)
	}

	// 返回用户信息和令牌
	// 注意：不要返回密码字段
	user.Password = ""

	return &models.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
