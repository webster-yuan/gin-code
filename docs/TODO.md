# 待办事项清单 (TODO List)

本文档用于记录项目的待办事项和功能改进计划。

---

## 🔐 认证授权相关

### RBAC 权限控制系统 ✅ 已完成

**优先级：** 低

**状态：** ✅ 已完成（2026-01-18）

**实现内容：**
1. ✅ **角色定义**：实现了 `RoleUser`（普通用户）和 `RoleAdmin`（超级管理员）两种角色
2. ✅ **用户角色关联**：在用户表中添加了 `role` 字段（INTEGER，默认值为0）
3. ✅ **权限中间件**：实现了 `RequireRole()` 和 `RequireAdmin()` 中间件
4. ✅ **JWT集成**：JWT Claims 中包含角色信息
5. ✅ **路由权限配置**：管理员路由（创建用户、删除用户）需要管理员权限
6. ✅ **国际化支持**：权限相关消息支持国际化

**实现文件：**
- `internal/auth/role.go` - 角色定义和权限检查方法
- `internal/auth/jwt.go` - JWT Claims 包含角色
- `internal/api/middleware/auth.go` - 权限检查中间件
- `internal/models/user.go` - 用户模型包含角色字段
- `internal/database/migrate.go` - 数据库迁移包含role字段
- `internal/repository/user_repository.go` - Repository层支持角色
- `internal/service/user_service.go` - Service层支持角色
- `internal/api/routes.go` - 路由权限配置
- `internal/i18n/i18n.go` - 权限相关国际化消息

**功能说明：**
详见 [RBAC权限控制功能说明.md](./RBAC权限控制功能说明.md)

**扩展性：**
- 代码结构支持添加新角色（使用 `iota`）
- `HasPermission()` 方法支持细粒度权限检查
- 数据库使用整数存储角色，便于扩展

---

## 📝 添加新待办事项

请在下方按格式添加新的待办事项：

### [功能/问题名称] ⚠️ [优先级]

**优先级：** 高/中/低

**问题描述：**
描述需要解决的问题或实现的功能。

**建议方案：**
描述建议的实现方案或解决思路。

**预期收益：**
列出实现后的预期收益。

**相关文件：**
列出相关的代码文件或文档。

---

**最后更新：** 2026-01-18
