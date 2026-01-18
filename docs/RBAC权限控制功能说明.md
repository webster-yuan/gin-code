# RBAC权限控制功能说明

## 概述

项目已实现基于角色的访问控制（Role-Based Access Control, RBAC）机制，支持超级管理员和普通用户两种角色，并保留了后续角色扩展的空间。

## 实现方案

### 技术栈

- **角色枚举**：使用 Go 的 `iota` 实现角色类型
- **JWT Claims**：角色信息存储在 JWT 令牌中
- **Gin中间件**：权限检查中间件
- **数据库字段**：用户表中存储角色信息

### 核心组件

| 组件 | 路径 | 功能 |
|------|------|------|
| 角色定义 | `internal/auth/role.go` | 角色类型定义和权限检查方法 |
| JWT Claims | `internal/auth/jwt.go` | JWT令牌中包含角色信息 |
| 权限中间件 | `internal/api/middleware/auth.go` | 角色权限检查中间件 |
| 用户模型 | `internal/models/user.go` | 用户模型包含角色字段 |
| 数据库迁移 | `internal/database/migrate.go` | 用户表包含role字段 |

## 角色定义

### 角色类型

```go
type Role int

const (
    RoleUser Role = iota  // 普通用户（0）
    RoleAdmin             // 超级管理员（1）
    // 后续可以扩展更多角色，例如：
    // RoleModerator  // 版主
    // RoleEditor     // 编辑
    // RoleViewer     // 查看者
)
```

### 角色方法

- **`String()`**：返回角色的字符串表示（"user" 或 "admin"）
- **`ParseRole(s string)`**：从字符串解析角色
- **`IsAdmin()`**：检查是否为管理员
- **`HasPermission(permission string)`**：检查角色是否有特定权限（可扩展）

## 数据库设计

### 用户表结构

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    age INTEGER NOT NULL DEFAULT 0,
    role INTEGER NOT NULL DEFAULT 0,  -- 角色：0=普通用户，1=超级管理员
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

**角色字段说明：**
- `role INTEGER NOT NULL DEFAULT 0`：角色字段，默认值为0（普通用户）
- 0 = 普通用户（RoleUser）
- 1 = 超级管理员（RoleAdmin）

## 权限控制流程

### 1. 用户注册

1. 用户提交注册信息
2. 系统创建用户，**默认角色为普通用户**（`RoleUser`）
3. 角色信息存储在数据库中

### 2. 用户登录

1. 用户提交邮箱和密码
2. 系统验证邮箱和密码
3. 验证成功后生成JWT令牌，**令牌中包含用户角色信息**
4. 返回访问令牌和刷新令牌

### 3. 访问受保护资源

1. 用户在请求头中携带JWT令牌：`Authorization: Bearer {token}`
2. **认证中间件**拦截请求，验证令牌并提取角色信息
3. 角色信息存储在请求上下文中：`c.Set("role", claims.Role)`
4. **权限中间件**检查用户角色是否符合要求
5. 权限不足返回403 Forbidden错误

## 权限中间件

### RequireRole 中间件

要求特定角色的中间件：

```go
// 要求管理员角色
users.Use(middleware.RequireAdmin())

// 要求特定角色
users.Use(middleware.RequireRole(auth.RoleUser))
```

**工作原理：**
1. 从请求上下文中获取用户角色
2. 检查角色是否存在
3. **管理员拥有所有权限**（`role.IsAdmin()` 返回 true）
4. 检查是否匹配所需角色
5. 权限不足返回403错误

### RequireAdmin 中间件

要求管理员角色的快捷方法：

```go
adminUsers.Use(middleware.RequireAdmin())
```

## 路由权限配置

### 当前权限配置

```go
// 用户相关路由（需要认证）
users := apiGroup.Group("/users")
users.Use(middleware.NewAuthMiddleware()) // 应用认证中间件
{
    // 需要管理员权限的路由
    adminUsers := users.Group("")
    adminUsers.Use(middleware.RequireAdmin()) // 应用管理员权限检查
    {
        adminUsers.POST("", userHandler.CreateUser())       // POST /api/v1/users（仅管理员）
        adminUsers.DELETE("/:id", userHandler.DeleteUser()) // DELETE /api/v1/users/:id（仅管理员）
    }

    // 普通用户和管理员都可以访问的路由
    users.GET("", userHandler.GetAllUsers())       // GET /api/v1/users
    users.GET("/:id", userHandler.GetUser())       // GET /api/v1/users/:id
    users.PUT("/:id", userHandler.UpdateUser())    // PUT /api/v1/users/:id
}
```

### 权限矩阵

| 路由 | 方法 | 普通用户 | 管理员 | 说明 |
|------|------|---------|--------|------|
| `/api/v1/users` | POST | ❌ | ✅ | 创建用户（仅管理员） |
| `/api/v1/users` | GET | ✅ | ✅ | 获取所有用户 |
| `/api/v1/users/:id` | GET | ✅ | ✅ | 获取单个用户 |
| `/api/v1/users/:id` | PUT | ✅ | ✅ | 更新用户 |
| `/api/v1/users/:id` | DELETE | ❌ | ✅ | 删除用户（仅管理员） |

## API接口

### 1. 用户注册（公开接口）

```
POST /api/v1/auth/register
Content-Type: application/json
```

**请求体：**

```json
{
  "name": "测试用户",
  "email": "test@example.com",
  "password": "123456",
  "age": 25
}
```

**响应：**

```json
{
  "code": 201,
  "message": "注册成功",
  "data": {
    "id": 1,
    "name": "测试用户",
    "email": "test@example.com",
    "age": 25,
    "role": 0,
    "created_at": "2026-01-18T15:00:00Z",
    "updated_at": "2026-01-18T15:00:00Z"
  }
}
```

**注意：** 注册的用户默认角色为 `0`（普通用户）

### 2. 用户登录（公开接口）

```
POST /api/v1/auth/login
Content-Type: application/json
```

**响应中的用户信息包含角色：**

```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "测试用户",
      "email": "test@example.com",
      "age": 25,
      "role": 0,
      "created_at": "2026-01-18T15:00:00Z",
      "updated_at": "2026-01-18T15:00:00Z"
    }
  }
}
```

### 3. 创建用户（需要管理员权限）

```
POST /api/v1/users
Content-Type: application/json
Authorization: Bearer {admin_token}
```

**权限要求：** 管理员（RoleAdmin）

**响应（权限不足）：**

```json
{
  "code": 403,
  "message": "权限不足",
  "timestamp": 1768662000,
  "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 4. 删除用户（需要管理员权限）

```
DELETE /api/v1/users/:id
Authorization: Bearer {admin_token}
```

**权限要求：** 管理员（RoleAdmin）

## 代码实现

### 角色定义

```go
// internal/auth/role.go
package auth

type Role int

const (
    RoleUser Role = iota  // 普通用户
    RoleAdmin             // 超级管理员
)

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

func (r Role) IsAdmin() bool {
    return r == RoleAdmin
}
```

### JWT Claims 包含角色

```go
// internal/auth/jwt.go
type UserClaims struct {
    UserID int64  `json:"user_id"`
    Email  string `json:"email"`
    Name   string `json:"name"`
    Role   Role   `json:"role"` // 用户角色
    jwt.RegisteredClaims
}
```

### 权限检查中间件

```go
// internal/api/middleware/auth.go
func RequireRole(requiredRole auth.Role) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 从上下文中获取用户角色
        roleValue, exists := c.Get("role")
        if !exists {
            response.Forbidden(c, i18n.UserMessage(i18n.UserPermissionDenied), nil)
            c.Abort()
            return
        }

        role, ok := roleValue.(auth.Role)
        if !ok {
            response.Forbidden(c, i18n.UserMessage(i18n.UserPermissionDenied), nil)
            c.Abort()
            return
        }

        // 管理员拥有所有权限
        if role.IsAdmin() {
            c.Next()
            return
        }

        // 检查是否匹配所需角色
        if role != requiredRole {
            response.Forbidden(c, i18n.UserMessage(i18n.UserPermissionDenied), nil)
            c.Abort()
            return
        }

        c.Next()
    }
}

// 要求管理员角色的快捷方法
func RequireAdmin() gin.HandlerFunc {
    return RequireRole(auth.RoleAdmin)
}
```

## 角色扩展

### 扩展新角色

如需添加新角色，只需在 `internal/auth/role.go` 中添加：

```go
const (
    RoleUser Role = iota
    RoleAdmin
    RoleModerator  // 新增：版主
    RoleEditor     // 新增：编辑
)
```

然后更新 `String()` 和 `ParseRole()` 方法：

```go
func (r Role) String() string {
    switch r {
    case RoleAdmin:
        return "admin"
    case RoleUser:
        return "user"
    case RoleModerator:
        return "moderator"
    case RoleEditor:
        return "editor"
    default:
        return "unknown"
    }
}
```

### 扩展权限检查

可以在 `HasPermission()` 方法中实现更细粒度的权限检查：

```go
func (r Role) HasPermission(permission string) bool {
    // 管理员拥有所有权限
    if r.IsAdmin() {
        return true
    }

    // 根据角色和权限进行判断
    switch r {
    case RoleModerator:
        return permission == "read" || permission == "moderate"
    case RoleEditor:
        return permission == "read" || permission == "write"
    case RoleUser:
        return permission == "read"
    default:
        return false
    }
}
```

## 错误处理

### 权限不足错误

当用户权限不足时，返回403 Forbidden错误：

```json
{
  "code": 403,
  "message": "权限不足",
  "timestamp": 1768662000,
  "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 错误日志

权限检查失败时，会记录日志（英文）：

```
Permission denied: insufficient role
request_id: xxx
path: /api/v1/users
method: POST
user_role: user
required_role: admin
```

## 国际化

### 日志消息（英文）

- `LogPermissionDenied`：权限不足：角色权限不够
- `LogPermissionDeniedNoRole`：权限不足：未找到角色信息
- `LogPermissionDeniedInvalidRole`：权限不足：无效的角色信息

### API响应消息（中文）

- `UserPermissionDenied`：权限不足

## 测试

### 测试管理员权限

1. **创建管理员用户**（需要在数据库中手动设置 `role = 1`）
2. **登录获取管理员令牌**
3. **访问管理员接口**（如 `POST /api/v1/users`）
4. **验证权限检查**

### 测试普通用户权限

1. **注册新用户**（默认角色为普通用户）
2. **登录获取用户令牌**
3. **尝试访问管理员接口**（应返回403错误）
4. **访问普通用户接口**（应成功）

## 最佳实践

### 1. 角色设计

- ✅ **最小权限原则**：默认角色为普通用户
- ✅ **可扩展性**：使用 `iota` 便于扩展新角色
- ✅ **类型安全**：使用枚举类型而非字符串

### 2. 权限检查

- ✅ **中间件方式**：在路由层进行权限检查
- ✅ **管理员特权**：管理员自动拥有所有权限
- ✅ **明确错误**：权限不足时返回明确的错误信息

### 3. 安全性

- ✅ **角色存储在JWT中**：无需每次查询数据库
- ✅ **角色验证**：在认证中间件中验证角色
- ✅ **日志记录**：记录权限检查失败的日志

### 4. 扩展性

- ✅ **预留扩展空间**：代码结构支持添加新角色
- ✅ **权限方法**：`HasPermission()` 方法支持细粒度权限
- ✅ **数据库兼容**：使用整数存储角色，便于扩展

## 注意事项

1. **默认角色**：新注册的用户默认为普通用户（`RoleUser`）
2. **管理员创建**：需要在数据库中手动将用户的 `role` 字段设置为 `1`
3. **角色更新**：可以通过更新用户接口修改用户角色（需要管理员权限）
4. **JWT令牌**：角色信息存储在JWT中，修改角色后需要重新登录获取新令牌
5. **权限缓存**：由于角色存储在JWT中，修改角色后需要等待令牌过期或重新登录

## 文件清单

### 核心文件

- `internal/auth/role.go` - 角色定义和权限检查
- `internal/auth/jwt.go` - JWT Claims包含角色
- `internal/api/middleware/auth.go` - 权限检查中间件
- `internal/models/user.go` - 用户模型包含角色字段
- `internal/database/migrate.go` - 数据库迁移包含role字段
- `internal/repository/user_repository.go` - Repository层支持角色
- `internal/service/user_service.go` - Service层支持角色
- `internal/api/routes.go` - 路由权限配置
- `internal/i18n/i18n.go` - 权限相关国际化消息

## 总结

RBAC权限控制功能实现了：

1. ✅ **角色定义**：支持普通用户和管理员两种角色
2. ✅ **权限检查**：基于角色的权限检查中间件
3. ✅ **JWT集成**：角色信息存储在JWT令牌中
4. ✅ **数据库支持**：用户表包含角色字段
5. ✅ **可扩展性**：代码结构支持添加新角色和权限
6. ✅ **国际化**：权限相关消息支持国际化
7. ✅ **错误处理**：权限不足时返回明确的错误信息

当前实现满足了企业级开发的基本需求，同时保留了后续扩展的空间。可以根据实际业务需求添加更多角色和更细粒度的权限控制。

---

**最后更新：** 2026-01-18
