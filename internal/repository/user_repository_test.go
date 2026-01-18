package repository

import (
	"context"
	"testing"

	"gin/internal/database"
	"gin/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestDB 创建测试数据库
func setupTestDB(t *testing.T) database.DB {
	// 使用内存数据库进行测试
	db, err := database.InitDB("sqlite3", ":memory:")
	require.NoError(t, err, "应该能创建测试数据库")

	// 初始化表结构
	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			age INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err = db.Exec(createTable)
	require.NoError(t, err, "应该能创建表")

	// 创建索引
	createIndex := `CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)`
	_, err = db.Exec(createIndex)
	require.NoError(t, err, "应该能创建索引")

	return db
}

// teardownTestDB 清理测试数据库
func teardownTestDB(t *testing.T, db database.DB) {
	if db != nil {
		err := db.Close()
		assert.NoError(t, err, "应该能关闭数据库")
	}
}

// TestUserRepository_Create 测试创建用户
func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("成功创建用户", func(t *testing.T) {
		user := &models.User{
			Name:     "张三",
			Email:    "zhangsan@example.com",
			Password: "hashed_password_for_test",
			Age:      25,
		}

		created, err := repo.Create(ctx, user)
		require.NoError(t, err, "应该能创建用户")
		assert.NotZero(t, created.ID, "用户ID应该不为0")
		assert.Equal(t, "张三", created.Name)
		assert.Equal(t, "zhangsan@example.com", created.Email)
		assert.Equal(t, 25, created.Age)
		assert.False(t, created.CreatedAt.IsZero())
		assert.False(t, created.UpdatedAt.IsZero())
	})

	t.Run("创建重复邮箱应该失败", func(t *testing.T) {
		user1 := &models.User{
			Name:     "用户1",
			Email:    "duplicate@example.com",
			Password: "hashed_password1",
			Age:      20,
		}
		_, err := repo.Create(ctx, user1)
		require.NoError(t, err, "第一次创建应该成功")

		user2 := &models.User{
			Name:     "用户2",
			Email:    "duplicate@example.com", // 重复邮箱
			Password: "hashed_password2",
			Age:      30,
		}
		_, err = repo.Create(ctx, user2)
		assert.Error(t, err, "重复邮箱应该失败")
	})
}

// TestUserRepository_FindByID 测试根据ID查找用户
func TestUserRepository_FindByID(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("成功查找存在的用户", func(t *testing.T) {
		// 先创建用户
		user := &models.User{
			Name:     "李四",
			Email:    "lisi@example.com",
			Password: "hashed_password",
			Age:      30,
		}
		created, err := repo.Create(ctx, user)
		require.NoError(t, err)

		// 查找用户
		found, err := repo.FindByID(ctx, created.ID)
		require.NoError(t, err)
		assert.Equal(t, created.ID, found.ID)
		assert.Equal(t, "李四", found.Name)
		assert.Equal(t, "lisi@example.com", found.Email)
		assert.Equal(t, 30, found.Age)
	})

	t.Run("查找不存在的用户应该失败", func(t *testing.T) {
		_, err := repo.FindByID(ctx, 99999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")
	})
}

// TestUserRepository_FindByEmail 测试根据邮箱查找用户
func TestUserRepository_FindByEmail(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("成功查找存在的用户", func(t *testing.T) {
		// 先创建用户
		user := &models.User{
			Name:     "王五",
			Email:    "wangwu@example.com",
			Password: "hashed_password",
			Age:      28,
		}
		created, err := repo.Create(ctx, user)
		require.NoError(t, err)

		// 根据邮箱查找
		found, err := repo.FindByEmail(ctx, "wangwu@example.com")
		require.NoError(t, err)
		assert.Equal(t, created.ID, found.ID)
		assert.Equal(t, "王五", found.Name)
		assert.Equal(t, "wangwu@example.com", found.Email)
	})

	t.Run("查找不存在的邮箱应该失败", func(t *testing.T) {
		_, err := repo.FindByEmail(ctx, "notexist@example.com")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")
	})
}

// TestUserRepository_FindAll 测试查找所有用户
func TestUserRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("空表应该返回空列表", func(t *testing.T) {
		users, err := repo.FindAll(ctx)
		require.NoError(t, err)
		assert.Empty(t, users)
	})

	t.Run("应该返回所有用户", func(t *testing.T) {
		// 创建多个用户
		users := []*models.User{
			{Name: "用户1", Email: "user1@example.com", Password: "hashed_password1", Age: 20},
			{Name: "用户2", Email: "user2@example.com", Password: "hashed_password2", Age: 25},
			{Name: "用户3", Email: "user3@example.com", Password: "hashed_password3", Age: 30},
		}

		for _, user := range users {
			_, err := repo.Create(ctx, user)
			require.NoError(t, err)
		}

		// 查找所有用户
		all, err := repo.FindAll(ctx)
		require.NoError(t, err)
		assert.Len(t, all, 3, "应该返回3个用户")
	})
}

// TestUserRepository_Update 测试更新用户
func TestUserRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("成功更新用户", func(t *testing.T) {
		// 先创建用户
		user := &models.User{
			Name:     "原始名称",
			Email:    "original@example.com",
			Password: "hashed_password",
			Age:      20,
		}
		created, err := repo.Create(ctx, user)
		require.NoError(t, err)

		// 更新用户
		updatedUser := &models.User{
			Name:     "更新后的名称",
			Email:    "updated@example.com",
			Password: "hashed_password",
			Age:      25,
		}
		updated, err := repo.Update(ctx, created.ID, updatedUser)
		require.NoError(t, err)
		assert.Equal(t, created.ID, updated.ID)
		assert.Equal(t, "更新后的名称", updated.Name)
		assert.Equal(t, "updated@example.com", updated.Email)
		assert.Equal(t, 25, updated.Age)
		// SQLite 的时间精度可能不够，只验证 UpdatedAt 不为零
		assert.False(t, updated.UpdatedAt.IsZero(), "更新时间应该设置")
	})

	t.Run("更新不存在的用户应该失败", func(t *testing.T) {
		user := &models.User{
			Name:     "测试",
			Email:    "test@example.com",
			Password: "hashed_password",
			Age:      20,
		}
		_, err := repo.Update(ctx, 99999, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")
	})
}

// TestUserRepository_Delete 测试删除用户
func TestUserRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("成功删除用户", func(t *testing.T) {
		// 先创建用户
		user := &models.User{
			Name:     "待删除用户",
			Email:    "todelete@example.com",
			Password: "hashed_password",
			Age:      20,
		}
		created, err := repo.Create(ctx, user)
		require.NoError(t, err)

		// 删除用户
		err = repo.Delete(ctx, created.ID)
		require.NoError(t, err)

		// 验证用户已删除
		_, err = repo.FindByID(ctx, created.ID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")
	})

	t.Run("删除不存在的用户应该失败", func(t *testing.T) {
		err := repo.Delete(ctx, 99999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户不存在")
	})
}

// TestUserRepository_Integration 集成测试：完整的CRUD流程
func TestUserRepository_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	repo := NewUserRepository(db)
	ctx := context.Background()

	// 1. 创建用户
	user := &models.User{
		Name:     "集成测试用户",
		Email:    "integration@example.com",
		Password: "hashed_password",
		Age:      25,
	}
	created, err := repo.Create(ctx, user)
	require.NoError(t, err)
	assert.NotZero(t, created.ID)

	// 2. 查找用户
	found, err := repo.FindByID(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)

	// 3. 更新用户
	updatedUser := &models.User{
		Name:     "更新后的集成测试用户",
		Email:    "integration@example.com",
		Password: "hashed_password",
		Age:      30,
	}
	updated, err := repo.Update(ctx, created.ID, updatedUser)
	require.NoError(t, err)
	assert.Equal(t, "更新后的集成测试用户", updated.Name)
	assert.Equal(t, 30, updated.Age)

	// 4. 删除用户
	err = repo.Delete(ctx, created.ID)
	require.NoError(t, err)

	// 5. 验证用户已删除
	_, err = repo.FindByID(ctx, created.ID)
	assert.Error(t, err)
}
