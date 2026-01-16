package service

import (
	"context"
	"errors"
	"testing"

	"gin/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockUserRepository 是 UserRepository 的 mock 实现
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, id int64, user *models.User) (*models.User, error) {
	args := m.Called(ctx, id, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// TestUserService_CreateUser 测试创建用户服务
func TestUserService_CreateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("成功创建用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		req := &models.CreateUserRequest{
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}

		// Mock: 邮箱不存在
		mockRepo.On("FindByEmail", ctx, req.Email).Return(nil, errors.New("用户不存在"))

		// Mock: 创建用户成功
		expectedUser := &models.User{
			ID:    1,
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}
		mockRepo.On("Create", ctx, mock.AnythingOfType("*models.User")).Return(expectedUser, nil)

		user, err := service.CreateUser(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, "张三", user.Name)
		assert.Equal(t, "zhangsan@example.com", user.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("邮箱已存在应该失败", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		req := &models.CreateUserRequest{
			Name:  "张三",
			Email: "existing@example.com",
			Age:   25,
		}

		// Mock: 邮箱已存在
		existingUser := &models.User{ID: 1, Email: "existing@example.com"}
		mockRepo.On("FindByEmail", ctx, req.Email).Return(existingUser, nil)

		user, err := service.CreateUser(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "邮箱已被使用")

		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "Create")
	})
}

// TestUserService_GetUserByID 测试根据ID获取用户
func TestUserService_GetUserByID(t *testing.T) {
	ctx := context.Background()

	t.Run("成功获取用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		expectedUser := &models.User{
			ID:    1,
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}

		mockRepo.On("FindByID", ctx, int64(1)).Return(expectedUser, nil)

		user, err := service.GetUserByID(ctx, 1)
		require.NoError(t, err)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, "张三", user.Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("无效的ID应该失败", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		user, err := service.GetUserByID(ctx, 0)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "用户ID无效")

		mockRepo.AssertNotCalled(t, "FindByID")
	})

	t.Run("用户不存在应该失败", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		mockRepo.On("FindByID", ctx, int64(999)).Return(nil, errors.New("用户不存在"))

		user, err := service.GetUserByID(ctx, 999)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "用户不存在")

		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_GetUserByEmail 测试根据邮箱获取用户
func TestUserService_GetUserByEmail(t *testing.T) {
	ctx := context.Background()

	t.Run("成功获取用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		expectedUser := &models.User{
			ID:    1,
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}

		mockRepo.On("FindByEmail", ctx, "zhangsan@example.com").Return(expectedUser, nil)

		user, err := service.GetUserByEmail(ctx, "zhangsan@example.com")
		require.NoError(t, err)
		assert.Equal(t, expectedUser.Email, user.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("空邮箱应该失败", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		user, err := service.GetUserByEmail(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "邮箱不能为空")

		mockRepo.AssertNotCalled(t, "FindByEmail")
	})
}

// TestUserService_GetAllUsers 测试获取所有用户
func TestUserService_GetAllUsers(t *testing.T) {
	ctx := context.Background()

	t.Run("成功获取所有用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		expectedUsers := []*models.User{
			{ID: 1, Name: "用户1", Email: "user1@example.com", Age: 20},
			{ID: 2, Name: "用户2", Email: "user2@example.com", Age: 25},
		}

		mockRepo.On("FindAll", ctx).Return(expectedUsers, nil)

		users, err := service.GetAllUsers(ctx)
		require.NoError(t, err)
		assert.Len(t, users, 2)

		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_UpdateUser 测试更新用户
func TestUserService_UpdateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("成功更新用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		existingUser := &models.User{
			ID:    1,
			Name:  "原始名称",
			Email: "original@example.com",
			Age:   20,
		}

		updatedUser := &models.User{
			ID:    1,
			Name:  "更新后的名称",
			Email: "original@example.com",
			Age:   25,
		}

		req := &models.UpdateUserRequest{
			Name: "更新后的名称",
			Age:  25,
		}

		mockRepo.On("FindByID", ctx, int64(1)).Return(existingUser, nil)
		mockRepo.On("Update", ctx, int64(1), mock.AnythingOfType("*models.User")).Return(updatedUser, nil)

		user, err := service.UpdateUser(ctx, 1, req)
		require.NoError(t, err)
		assert.Equal(t, "更新后的名称", user.Name)
		assert.Equal(t, 25, user.Age)

		mockRepo.AssertExpectations(t)
	})

	t.Run("更新邮箱时检查重复", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		existingUser := &models.User{
			ID:    1,
			Name:  "用户1",
			Email: "user1@example.com",
			Age:   20,
		}

		req := &models.UpdateUserRequest{
			Email: "existing@example.com",
		}

		// 用户存在
		mockRepo.On("FindByID", ctx, int64(1)).Return(existingUser, nil)
		// 新邮箱已存在
		existingUser2 := &models.User{ID: 2, Email: "existing@example.com"}
		mockRepo.On("FindByEmail", ctx, "existing@example.com").Return(existingUser2, nil)

		user, err := service.UpdateUser(ctx, 1, req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "邮箱已被使用")

		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "Update")
	})
}

// TestUserService_DeleteUser 测试删除用户
func TestUserService_DeleteUser(t *testing.T) {
	ctx := context.Background()

	t.Run("成功删除用户", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		existingUser := &models.User{
			ID:    1,
			Name:  "待删除用户",
			Email: "todelete@example.com",
			Age:   20,
		}

		mockRepo.On("FindByID", ctx, int64(1)).Return(existingUser, nil)
		mockRepo.On("Delete", ctx, int64(1)).Return(nil)

		err := service.DeleteUser(ctx, 1)
		require.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("删除不存在的用户应该失败", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		mockRepo.On("FindByID", ctx, int64(999)).Return(nil, errors.New("用户不存在"))

		err := service.DeleteUser(ctx, 999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")

		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "Delete")
	})
}
