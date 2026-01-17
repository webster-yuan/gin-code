package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gin/internal/database"
	"gin/internal/models"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, id int64, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int64) error
}

// userRepository 用户仓库实现
type userRepository struct {
	db database.DB
}

// NewUserRepository 创建用户仓库
func NewUserRepository(db database.DB) UserRepository {
	return &userRepository{db: db}
}

// Create 创建用户
func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// SQLite 不支持 RETURNING，使用 Exec + LastInsertId
	result, err := r.db.Exec(
		"INSERT INTO users (name, email, password, age, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Age, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("获取用户ID失败: %w", err)
	}

	// 查询刚创建的用户
	return r.FindByID(ctx, id)
}

// FindByID 根据ID查找用户
func (r *userRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, name, email, age, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("用户不存在: %w", err)
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password, age, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("用户不存在: %w", err)
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return user, nil
}

// FindAll 查找所有用户
func (r *userRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	query := `
		SELECT id, name, email, age, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询用户列表失败: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描用户数据失败: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历用户数据失败: %w", err)
	}

	return users, nil
}

// Update 更新用户
func (r *userRepository) Update(ctx context.Context, id int64, user *models.User) (*models.User, error) {
	user.UpdatedAt = time.Now()

	query := `
		UPDATE users
		SET name = ?, email = ?, age = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := r.db.Exec(query, user.Name, user.Email, user.Age, user.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("获取影响行数失败: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}

	return r.FindByID(ctx, id)
}

// Delete 删除用户
func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}
