# Gin Web框架示例项目

这是一个基于Gin框架的Go Web应用程序示例，展示了Gin的各种功能和最佳实践。

## 项目结构
- `cmd/server/` - 服务器启动入口
- `internal/` - 内部包，包含应用核心代码
  - `api/` - API路由和处理
  - `middleware/` - 中间件
  - `models/` - 数据模型
- `pkg/` - 可重用的公共包
- `data_structures/` - 数据结构实现
- `examples/` - 示例代码
- `templates/` - HTML模板文件
- `static/` - 静态资源文件

## 功能特性
- 多服务器并发启动
- 路由分组和中间件
- 模板渲染
- 文件上传
- 参数绑定和验证
- JSON处理
- Protocol Buffers支持
- 重定向功能

## 快速开始
1. 确保安装了Go 1.24.0或更高版本
2. 克隆仓库
3. 运行 `go mod tidy` 安装依赖
4. 运行 `go run main.go` 启动应用
5. 访问 http://localhost:8080/hello 查看基本响应

## 主要端口
- 主服务: 8080
- 额外服务1: 8081
- 额外服务2: 8082