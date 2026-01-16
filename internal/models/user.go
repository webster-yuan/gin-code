package models

import "time"

// User 用户模型
type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required,min=2,max=50"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Age       int       `json:"age" db:"age" binding:"gte=0,lte=150"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=50"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=0,lte=150"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=50"`
	Email string `json:"email" binding:"omitempty,email"`
	Age   int    `json:"age" binding:"omitempty,gte=0,lte=150"`
}
