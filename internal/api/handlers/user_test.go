package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin/internal/errors"
	"gin/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockUserService 是 UserService 的 mock 实现
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id int64, req *models.UpdateUserRequest) (*models.User, error) {
	args := m.Called(ctx, id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.LoginResponse), args.Error(1)
}

// setupTestRouter 设置测试路由
func setupTestRouter(handler *UserHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 添加错误处理中间件（与生产环境一致）
	router.Use(errors.ErrorHandler())

	api := router.Group("/api/v1")
	users := api.Group("/users")
	{
		users.POST("", handler.CreateUser())
		users.GET("", handler.GetAllUsers())
		users.GET("/:id", handler.GetUser())
		users.PUT("/:id", handler.UpdateUser())
		users.DELETE("/:id", handler.DeleteUser())
	}

	return router
}

// TestUserHandler_CreateUser 测试创建用户处理器
func TestUserHandler_CreateUser(t *testing.T) {
	t.Run("成功创建用户", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		reqBody := models.CreateUserRequest{
			Name:     "张三",
			Email:    "zhangsan@example.com",
			Password: "123456",
			Age:      25,
		}

		expectedUser := &models.User{
			ID:    1,
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}

		mockService.On("CreateUser", mock.Anything, &reqBody).Return(expectedUser, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, float64(201), response["code"])
		assert.Equal(t, "创建成功", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("无效的请求体应该返回400", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertNotCalled(t, "CreateUser")
	})
}

// TestUserHandler_GetUser 测试获取用户处理器
func TestUserHandler_GetUser(t *testing.T) {
	t.Run("成功获取用户", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		expectedUser := &models.User{
			ID:    1,
			Name:  "张三",
			Email: "zhangsan@example.com",
			Age:   25,
		}

		mockService.On("GetUserByID", mock.Anything, int64(1)).Return(expectedUser, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, float64(200), response["code"])

		mockService.AssertExpectations(t)
	})

	t.Run("无效的ID应该返回400", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/users/invalid", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertNotCalled(t, "GetUserByID")
	})
}

// TestUserHandler_GetAllUsers 测试获取所有用户处理器
func TestUserHandler_GetAllUsers(t *testing.T) {
	t.Run("成功获取所有用户", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		expectedUsers := []*models.User{
			{ID: 1, Name: "用户1", Email: "user1@example.com", Age: 20},
			{ID: 2, Name: "用户2", Email: "user2@example.com", Age: 25},
		}

		mockService.On("GetAllUsers", mock.Anything).Return(expectedUsers, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, float64(200), response["code"])

		mockService.AssertExpectations(t)
	})
}

// TestUserHandler_UpdateUser 测试更新用户处理器
func TestUserHandler_UpdateUser(t *testing.T) {
	t.Run("成功更新用户", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		reqBody := models.UpdateUserRequest{
			Name: "更新后的名称",
			Age:  30,
		}

		updatedUser := &models.User{
			ID:    1,
			Name:  "更新后的名称",
			Email: "zhangsan@example.com",
			Age:   30,
		}

		mockService.On("UpdateUser", mock.Anything, int64(1), &reqBody).Return(updatedUser, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/users/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "更新成功", response["message"])

		mockService.AssertExpectations(t)
	})
}

// TestUserHandler_DeleteUser 测试删除用户处理器
func TestUserHandler_DeleteUser(t *testing.T) {
	t.Run("成功删除用户", func(t *testing.T) {
		mockService := new(MockUserService)
		handler := NewUserHandler(mockService)
		router := setupTestRouter(handler)

		mockService.On("DeleteUser", mock.Anything, int64(1)).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/1", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "删除成功", response["message"])

		mockService.AssertExpectations(t)
	})
}
