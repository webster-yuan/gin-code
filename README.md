# Go Startup - Gin Web 开发学习项目

一个完整的 Gin Web 开发学习项目，涵盖企业级开发的核心步骤和实践。

**项目状态：** ✅ **基本框架已完成（90%）**，可用于企业级项目开发和学习

## ✨ 项目特性

### 核心架构
- ✅ **三层架构**：Handler → Service → Repository → Database
- ✅ **依赖注入**：DI 容器统一管理依赖
- ✅ **接口抽象**：Repository 和 Service 使用接口，便于测试和扩展
- ✅ **项目结构**：符合 Go 项目标准布局

### 配置和日志
- ✅ **Viper 配置管理**：支持 YAML、环境变量
- ✅ **Zap 结构化日志**：高性能日志库，支持国际化（英文日志）
- ✅ **统一错误处理**：AppError 结构体，统一错误响应格式

### 安全认证
- ✅ **JWT 认证**：基于 JWT 的认证机制
- ✅ **刷新令牌机制**：访问令牌和刷新令牌分离
- ✅ **密码加密**：使用 bcrypt 哈希密码
- ✅ **RBAC 权限控制**：支持超级管理员和普通用户两种角色
- ✅ **认证中间件**：JWT 验证和权限检查

### API 开发
- ✅ **统一响应格式**：所有 API 使用统一的响应结构
- ✅ **Swagger 文档**：自动生成 API 文档（`/swagger/index.html`）
- ✅ **API 版本控制**：`/api/v1` 路由前缀
- ✅ **数据验证**：使用 validator 标签验证请求参数
- ✅ **国际化**：日志英文、API响应中文

### 中间件体系
- ✅ **请求日志中间件**：记录请求和响应
- ✅ **错误处理中间件**：统一错误响应
- ✅ **异常恢复中间件**：防止 panic
- ✅ **Prometheus 指标**：监控指标收集
- ✅ **请求ID中间件**：为每个请求生成唯一ID，便于追踪
- ✅ **认证授权中间件**：JWT 验证和权限检查

### 数据库
- ✅ **SQLite 支持**：本地文件数据库（默认）
- ✅ **数据库抽象**：支持多种数据库驱动（MySQL、SQLite）
- ✅ **自动初始化**：首次启动自动创建表结构
- ✅ **连接池配置**：优化数据库连接

### 监控和可观测性
- ✅ **Prometheus 指标**：请求计数、响应时间统计（`/metrics`）
- ✅ **性能分析（pprof）**：CPU、内存分析（`/debug/pprof/*`）
- ✅ **健康检查**：服务状态监控（`/health`）

### 测试
- ✅ **单元测试**：Repository、Service、Handler 层全覆盖
- ✅ **测试覆盖率**：Repository 84.1%，Service 88.9%
- ✅ **Mock 支持**：使用 testify/mock

### 部署
- ✅ **优雅关闭**：信号监听、资源清理
- ✅ **Docker 支持**：容器化部署（Dockerfile）
- ✅ **多服务器支持**：使用 errgroup 并发启动

## 🚀 快速开始

### 环境要求

- Go 1.19+
- SQLite（项目自带，无需额外安装）

### 安装依赖

```bash
go mod download
```

### 配置

编辑 `internal/config/config.yaml`：

```yaml
server:
  port: "8080"
  read_timeout: 5
  write_timeout: 5

database:
  driver: "sqlite3"
  dsn: "./data/app.db"

logging:
  level: "info"

jwt:
  secret_key: "your-secret-key-change-in-production"
  expires_in: 24        # 访问令牌过期时间（小时）
  refresh_expires_in: 168  # 刷新令牌过期时间（小时，默认7天）
```

### 运行

```bash
# 方式1：直接运行
go run main.go server

# 方式2：编译后运行
go build -o bin/server.exe ./cmd/server
./bin/server.exe server
```

### 访问服务

- **API 服务**：http://localhost:8080
- **Swagger 文档**：http://localhost:8080/swagger/index.html
- **健康检查**：http://localhost:8080/health
- **Prometheus 指标**：http://localhost:8080/metrics
- **性能分析**：http://localhost:6060/debug/pprof/

## 📁 项目结构

```
.
├── cmd/                    # 命令行入口
│   ├── commands/          # 子命令
│   └── server/            # 服务器启动
├── internal/              # 内部代码
│   ├── api/               # API 层
│   │   ├── handlers/      # 处理器
│   │   ├── middleware/    # 中间件（认证、权限、请求ID）
│   │   ├── response/      # 统一响应格式
│   │   └── routes.go      # 路由配置
│   ├── auth/              # 认证授权
│   │   ├── jwt.go         # JWT 工具
│   │   ├── password.go    # 密码加密
│   │   └── role.go        # 角色定义
│   ├── service/           # 业务逻辑层
│   ├── repository/        # 数据访问层
│   ├── models/            # 数据模型
│   ├── config/            # 配置管理
│   ├── database/          # 数据库连接和迁移
│   ├── di/                # 依赖注入容器
│   ├── errors/            # 错误处理
│   ├── i18n/              # 国际化
│   ├── logger/            # 日志系统
│   ├── metrics/           # 指标监控
│   └── middleware/        # 应用中间件（Recovery）
├── docs/                  # 文档目录
│   ├── API文档配置说明.md
│   ├── RBAC权限控制功能说明.md
│   ├── 认证授权功能说明.md
│   ├── 国际化功能说明.md
│   ├── 统一响应格式说明.md
│   ├── 用户注册功能说明.md
│   └── ...
├── scripts/               # 脚本工具
└── data/                  # 数据文件
```

## 📊 API 端点

### 认证相关（公开接口）

- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录（返回 access_token 和 refresh_token）
- `POST /api/v1/auth/refresh` - 刷新访问令牌

### 用户相关（需要认证）

- `POST /api/v1/users` - 创建用户（**需要管理员权限**）
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取单个用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户（**需要管理员权限**）

### 监控端点

- `GET /health` - 健康检查
- `GET /metrics` - Prometheus 指标
- `GET /debug/pprof/*` - 性能分析端点
- `GET /swagger/*` - Swagger API 文档

## 📚 文档

### 核心文档

- [项目技术栈详解](./docs/项目技术栈详解.md) - 技术栈详细说明
- [框架完成度与环境准备说明](./docs/框架完成度与环境准备说明.md) - 框架完成度和环境准备指南
- [三层架构使用说明](./docs/三层架构使用说明.md) - 架构使用指南
- [单元测试说明](./docs/单元测试说明.md) - 测试编写指南

### 功能文档

- [认证授权功能说明](./docs/认证授权功能说明.md) - JWT认证和刷新令牌机制
- [RBAC权限控制功能说明](./docs/RBAC权限控制功能说明.md) - 角色权限控制
- [统一响应格式说明](./docs/统一响应格式说明.md) - API响应格式规范
- [国际化功能说明](./docs/国际化功能说明.md) - 国际化消息管理
- [用户注册功能说明](./docs/用户注册功能说明.md) - 用户注册流程

### 工具文档

- [Go性能分析指南](./docs/Go性能分析指南.md) - pprof 性能分析教程
- [API文档配置说明](./docs/API文档配置说明.md) - Swagger 配置指南

### 其他文档

- [TODO.md](./docs/TODO.md) - 待办事项清单

## 🛠️ 开发工具

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/service/... -v

# 查看测试覆盖率
go test ./internal/... -cover
```

### 性能分析

```bash
# 启动服务器后，使用脚本（Windows PowerShell）
.\scripts\profile.ps1

# 或直接使用命令
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```

### 生成 Swagger 文档

```bash
go run github.com/swaggo/swag/cmd/swag init -o internal/docs --parseDependency --parseInternal
```

## 🎯 学习路径

1. **基础架构** - 了解项目结构和三层架构
2. **配置管理** - 学习 Viper 配置管理
3. **日志系统** - 掌握 Zap 日志使用和国际化
4. **数据库操作** - Repository 层数据访问
5. **业务逻辑** - Service 层业务处理
6. **API 开发** - Handler 层接口开发
7. **安全认证** - JWT 认证和 RBAC 权限控制
8. **中间件开发** - 自定义中间件编写
9. **测试编写** - 单元测试和 Mock
10. **性能分析** - pprof 性能分析
11. **API 文档** - Swagger 文档生成和使用

## 📦 技术栈

### 核心框架

- **Web 框架**：Gin - 高性能 HTTP Web 框架
- **命令行框架**：Cobra - 强大的 CLI 应用框架
- **配置管理**：Viper - 支持多种配置格式

### 认证授权

- **JWT**：golang-jwt/jwt - JWT 令牌生成和验证
- **密码加密**：golang.org/x/crypto/bcrypt - 密码哈希

### 日志和监控

- **日志库**：Zap - 高性能结构化日志
- **监控**：Prometheus - 指标收集和监控
- **性能分析**：pprof - CPU、内存性能分析

### 数据库

- **数据库驱动**：SQLite（默认）、MySQL（支持）
- **接口抽象**：database/sql - 标准数据库接口

### 测试

- **测试框架**：testing - Go 标准测试库
- **断言库**：testify/assert - 测试断言
- **Mock 库**：testify/mock - Mock 对象

### 文档和工具

- **API 文档**：Swagger/OpenAPI - 自动生成 API 文档
- **国际化**：自定义 i18n 包 - 多语言消息管理

## 🔄 可扩展功能

以下功能可以根据实际需求按需添加：

### 性能优化（中优先级）

- [ ] **缓存层（Redis）** - 缓存热点数据，提升查询性能
- [ ] **限流中间件** - 防止接口被刷，保护系统资源
- [ ] **事务管理** - 保证数据一致性（已预留接口）

### 工程化工具（低优先级）

- [ ] **数据库迁移工具** - 版本化管理表结构变更
- [ ] **集成测试** - 端到端测试
- [ ] **CI/CD 配置** - 自动化构建和部署
- [ ] **CORS 配置** - 跨域请求支持

## 💻 环境准备

### 开发环境（Win10）

**当前项目可以直接在 Win10 上运行，无需额外环境！**

- ✅ **SQLite**：项目自带，无需安装
- ⬜ **Docker Desktop**（可选）：用于学习中间件（Redis、MySQL）
- ⬜ **Linux/K8s**：**不需要**，开发阶段完全不需要

详见：[框架完成度与环境准备说明](./docs/框架完成度与环境准备说明.md)

## 📈 项目状态

**框架完成度：约 90%**

✅ **已完成：**
- 核心架构和基础功能
- 安全认证和权限控制
- API 开发和文档
- 监控和测试

⚠️ **可扩展：**
- 中间件集成（Redis、MySQL 等，按需添加）
- 性能优化功能（缓存、限流、事务）
- 工程化工具（CI/CD、迁移工具）

## 🎓 适用场景

### ✅ 适合用于

- **学习企业级开发**：完整的企业级架构和实践
- **中小型项目开发**：框架已具备核心功能
- **API 服务开发**：完整的 RESTful API 开发能力
- **认证授权学习**：JWT、RBAC 完整实现

### ⚠️ 需要注意

- **缓存**：如需高并发，建议添加 Redis 缓存
- **限流**：如需防止接口被刷，建议添加限流中间件
- **事务**：复杂业务场景需要事务支持

## 📝 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

---

**最后更新：** 2026-01-18
