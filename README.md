# Go Startup - Gin Web 开发学习项目

一个完整的 Gin Web 开发学习项目，涵盖企业级开发的核心步骤和实践。

## ✨ 项目特性

### 核心架构
- ✅ **三层架构**：Handler → Service → Repository → Database
- ✅ **依赖注入**：统一的依赖管理
- ✅ **配置管理**：Viper 支持 YAML 和环境变量
- ✅ **日志系统**：Zap 结构化日志
- ✅ **错误处理**：统一的错误处理和响应

### 数据库
- ✅ **SQLite 支持**：本地文件数据库
- ✅ **数据库抽象**：支持多种数据库驱动
- ✅ **自动初始化**：首次启动自动创建表结构

### 监控和可观测性
- ✅ **Prometheus 指标**：请求计数、响应时间统计
- ✅ **性能分析**：pprof 支持 CPU、内存分析
- ✅ **健康检查**：服务状态监控

### 测试
- ✅ **单元测试**：Repository、Service、Handler 层全覆盖
- ✅ **测试覆盖率**：Repository 84.1%，Service 88.9%
- ✅ **Mock 支持**：使用 testify/mock

### 部署
- ✅ **优雅关闭**：信号监听、资源清理
- ✅ **Docker 支持**：容器化部署
- ✅ **多服务器**：支持并发启动多个服务

## 🚀 快速开始

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
```

### 运行

```bash
# 方式1：直接运行
go run main.go server

# 方式2：编译后运行
go build -o bin/server.exe ./cmd/server
./bin/server.exe server
```

### 测试 API

```bash
# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }'

# 获取所有用户
curl -X GET http://localhost:8080/api/v1/users
```

## 📁 项目结构

```
.
├── cmd/                    # 命令行入口
│   ├── commands/          # 子命令
│   └── server/            # 服务器启动
├── internal/              # 内部代码
│   ├── api/               # API 层
│   │   ├── handlers/      # 处理器
│   │   └── middleware/    # 中间件
│   ├── service/           # 业务逻辑层
│   ├── repository/        # 数据访问层
│   ├── models/            # 数据模型
│   ├── config/            # 配置管理
│   ├── database/          # 数据库
│   ├── errors/            # 错误处理
│   ├── logger/            # 日志系统
│   └── metrics/           # 指标监控
├── docs/                  # 文档
├── scripts/               # 脚本工具
└── data/                  # 数据文件
```

## 📚 文档

- [项目技术栈详解](./docs/项目技术栈详解.md) - 技术栈详细说明
- [企业级开发对比分析](./docs/企业级开发对比分析.md) - 与企业级项目对比
- [三层架构使用说明](./docs/三层架构使用说明.md) - 架构使用指南
- [单元测试说明](./docs/单元测试说明.md) - 测试编写指南
- [Go 性能分析指南](./docs/Go性能分析指南.md) - 性能分析教程
- [项目完成度评估](./docs/项目完成度评估.md) - 功能完成度评估

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
# 启动服务器后，使用脚本
.\scripts\profile.ps1

# 或直接使用命令
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```

## 📊 API 端点

### 用户相关 API

- `POST /api/v1/users` - 创建用户
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取单个用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 监控端点

- `GET /metrics` - Prometheus 指标
- `GET /debug/pprof/*` - 性能分析端点

## 🎯 学习路径

1. **基础架构** - 了解项目结构和三层架构
2. **配置管理** - 学习 Viper 配置管理
3. **日志系统** - 掌握 Zap 日志使用
4. **数据库操作** - Repository 层数据访问
5. **业务逻辑** - Service 层业务处理
6. **API 开发** - Handler 层接口开发
7. **测试编写** - 单元测试和 Mock
8. **性能分析** - pprof 性能分析

## 📦 技术栈

- **Web 框架**：Gin
- **配置管理**：Viper
- **日志库**：Zap
- **数据库**：SQLite (支持 MySQL)
- **测试框架**：testing + testify
- **监控**：Prometheus
- **性能分析**：pprof
- **命令行**：Cobra

## 🔄 待实现功能

- [ ] API 文档（Swagger）
- [ ] JWT 认证授权
- [ ] 统一响应格式
- [ ] 缓存层（Redis）
- [ ] 限流中间件
- [ ] 数据库迁移工具
- [ ] CI/CD 配置

详见：[项目完成度评估](./docs/项目完成度评估.md)

## 📝 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！
