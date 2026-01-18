package auth

// Role 角色类型
type Role int

const (
	// RoleUser 普通用户
	RoleUser Role = iota
	// RoleAdmin 超级管理员
	RoleAdmin
	// 后续可以扩展更多角色，例如：
	// RoleModerator  // 版主
	// RoleEditor     // 编辑
	// RoleViewer     // 查看者
)

// String 返回角色的字符串表示
func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	default:
		return "unknown"
	}
}

// ParseRole 从字符串解析角色
func ParseRole(s string) Role {
	switch s {
	case "admin":
		return RoleAdmin
	case "user":
		return RoleUser
	default:
		return RoleUser // 默认为普通用户
	}
}

// IsAdmin 检查是否为管理员
func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

// HasPermission 检查角色是否有权限（可扩展用于更复杂的权限检查）
func (r Role) HasPermission(permission string) bool {
	// 管理员拥有所有权限
	if r.IsAdmin() {
		return true
	}

	// 普通用户的权限检查（可根据业务需求扩展）
	switch permission {
	case "read":
		return true
	case "write":
		return false // 普通用户默认无写权限
	default:
		return false
	}
}
