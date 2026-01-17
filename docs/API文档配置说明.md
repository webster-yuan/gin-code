# API文档配置说明

## 概述

项目已集成Swagger/OpenAPI自动生成API文档功能，所有API接口都将自动生成可视化文档，便于开发人员和测试人员使用和调试API。

## 实现方式

使用以下技术栈实现API文档自动生成：

- **Swaggo/swag**：Go语言的Swagger文档生成工具
- **Swaggo/gin-swagger**：Gin框架的Swagger中间件
- **Swaggo/files**：Swagger UI的静态文件服务

## 目录结构

```
├── docs/                  # Markdown文档目录
├── internal/
│   └── docs/              # 自动生成的Swagger Go代码
└── main.go                # 项目入口，包含API基本信息注释
```

## 配置步骤

### 2. 安装依赖

```bash
go get -u github.com/swaggo/swag/cmd/swag github.com/swaggo/gin-swagger github.com/swaggo/files
```

### 3. 添加API基本信息注释

在`main.go`文件中添加API文档的基本信息：

```go
// Package main 是项目的主入口
// @title Go Startup API
// @version 1.0
// @description 这是一个基于Gin框架的企业级Go项目API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main
```

### 4. 为API端点添加注释

在每个API处理函数前添加Swagger注释：

```go
// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建一个新的用户记录
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "用户信息"
// @Success 201 {object} response.Response{data=models.User} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser() gin.HandlerFunc {
    // 处理逻辑...
}
```

### 5. 生成Swagger文档

由于swag命令可能不在系统路径中，我们使用go run命令执行：

```bash
go run github.com/swaggo/swag/cmd/swag init -o internal/docs
```

该命令会生成以下文件：
- `internal/docs/docs.go`：Go语言版本的Swagger文档
- `internal/docs/swagger.json`：JSON格式的Swagger文档
- `internal/docs/swagger.yaml`：YAML格式的Swagger文档

### 6. 配置路由

在Gin路由配置中添加Swagger UI的访问路径：

```go
import (
    _ "gin/internal/docs" // 导入Swagger文档
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    
    // 添加Swagger路由
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    
    // 其他路由配置...
    
    return router
}
```

## 使用方式

### 访问API文档

启动服务后，通过以下地址访问API文档：

```
http://localhost:8080/swagger/index.html
```

### 在线测试API

在Swagger UI中可以直接测试API：

1. 找到要测试的API端点
2. 点击"Try it out"
3. 填写请求参数
4. 点击"Execute"
5. 查看响应结果

## 注释规范

### API基本信息注释

| 注释标签 | 说明 | 示例 |
|---------|------|------|
| @title | API标题 | @title Go Startup API |
| @version | API版本 | @version 1.0 |
| @description | API描述 | @description 这是一个基于Gin框架的企业级Go项目API |
| @termsOfService | 服务条款 | @termsOfService http://swagger.io/terms/ |
| @contact.name | 联系人姓名 | @contact.name API Support |
| @contact.url | 联系人URL | @contact.url http://www.example.com/support |
| @contact.email | 联系人邮箱 | @contact.email support@example.com |
| @license.name | 许可证名称 | @license.name Apache 2.0 |
| @license.url | 许可证URL | @license.url http://www.apache.org/licenses/LICENSE-2.0.html |
| @host | API主机 | @host localhost:8080 |
| @BasePath | API基础路径 | @BasePath / |
| @schemes | API协议 | @schemes http |

### API端点注释

| 注释标签 | 说明 | 示例 |
|---------|------|------|
| @Summary | API摘要 | @Summary 创建新用户 |
| @Description | API详细描述 | @Description 创建一个新的用户记录 |
| @Tags | API标签（用于分组） | @Tags users |
| @Accept | 支持的请求格式 | @Accept json |
| @Produce | 支持的响应格式 | @Produce json |
| @Param | 请求参数 | @Param user body models.CreateUserRequest true "用户信息" |
| @Success | 成功响应 | @Success 201 {object} response.Response{data=models.User} "创建成功" |
| @Failure | 错误响应 | @Failure 400 {object} response.Response "请求参数错误" |
| @Router | API路由 | @Router /api/v1/users [post] |

## 示例API注释

### 创建用户

```go
// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建一个新的用户记录
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "用户信息"
// @Success 201 {object} response.Response{data=models.User} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var req models.CreateUserRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.Error(err)
            return
        }

        user, err := h.userService.CreateUser(c.Request.Context(), &req)
        if err != nil {
            c.Error(err)
            return
        }

        response.Created(c, "创建成功", user)
    }
}
```

### 获取用户

```go
// GetUser 获取用户
// @Summary 获取单个用户
// @Description 根据用户ID获取用户详细信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=models.User} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "用户不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            c.Error(err)
            return
        }

        user, err := h.userService.GetUserByID(c.Request.Context(), id)
        if err != nil {
            c.Error(err)
            return
        }

        response.Success(c, "获取成功", user)
    }
}
```

## 优势

1. **自动生成**：基于代码注释自动生成API文档
2. **可视化**：提供交互式的Swagger UI
3. **实时更新**：修改代码注释后重新生成即可更新文档
4. **标准化**：遵循OpenAPI规范
5. **易于测试**：支持在线测试API
6. **类型安全**：基于Go结构体自动生成参数和响应模型

## 注意事项

1. **注释格式**：严格按照Swagger注释规范编写
2. **代码位置**：自动生成的Go代码应放在internal/docs目录下，与Markdown文档分离
3. **定期更新**：修改API后应重新生成文档
4. **参数类型**：确保注释中的参数类型与实际代码一致
5. **响应结构**：使用统一响应格式的结构体

## 开发规范

### 文档目录结构

按照Gin框架的开发规范，文档目录应分为：

- `docs/`：存放Markdown文档（说明文档、设计文档等）
- `internal/docs/`：存放自动生成的Swagger Go代码

### 代码规范

1. **导入路径**：使用相对路径导入Swagger文档
2. **注释风格**：使用标准Go注释风格
3. **标签命名**：使用驼峰命名法
4. **参数描述**：清晰、简洁地描述参数用途
5. **响应示例**：提供完整的响应示例

## 迁移指南

### 生成新文档

```bash
# 清理旧文档
rm -rf internal/docs/

# 生成新文档
go run github.com/swaggo/swag/cmd/swag init -o internal/docs
```

### 更新路由配置

```go
// 旧导入
import (
    _ "gin/docs" // 旧路径
)

// 新导入
import (
    _ "gin/internal/docs" // 新路径
)
```

## 常见问题

### 文档生成失败

**问题**：`swag init` 命令失败

**解决方案**：
1. 检查代码注释格式是否正确
2. 确保所有引用的类型都已定义
3. 检查导入路径是否正确

### 文档无法访问

**问题**：访问`/swagger/index.html`显示404

**解决方案**：
1. 检查路由配置是否正确
2. 确保已导入Swagger文档包
3. 检查文档是否已生成

### 文档内容不完整

**问题**：生成的文档缺少API端点

**解决方案**：
1. 检查API处理函数是否添加了Swagger注释
2. 确保注释格式正确
3. 重新生成文档

## 总结

通过集成Swagger/OpenAPI，项目实现了API文档的自动生成和可视化，提高了开发效率和API的可维护性。遵循文档中描述的规范和最佳实践，可以确保API文档的质量和一致性。