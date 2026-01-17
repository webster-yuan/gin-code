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

import (
	"gin/cmd"
)

func main() {
	// 使用cobra的根命令
	cmd.Execute()
}
