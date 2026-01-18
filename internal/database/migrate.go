package database

import (
	"context"
	"fmt"
)

// InitSchema 初始化数据库表结构
func InitSchema(db DB) error {
	ctx := context.Background()

	// 创建 users 表
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			age INTEGER NOT NULL DEFAULT 0,
			role INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("创建 users 表失败: %w", err)
	}

	// 创建索引
	createIndex := `
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)
	`
	_, err = db.Exec(createIndex)
	if err != nil {
		return fmt.Errorf("创建索引失败: %w", err)
	}

	// 测试连接
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	return nil
}
