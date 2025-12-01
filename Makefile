.PHONY: build run test clean lint

# 设置Go环境变量
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

# 构建应用
build:
	@echo "Building application..."
	@go build -o output/gin-app main.go

# 运行应用
run:
	@echo "Running application..."
	@go run main.go server

# 运行测试
test:
	@echo "Running tests..."
	@go test ./... -v

# 代码格式化
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# 代码检查
lint:
	@echo "Linting code..."
	@golangci-lint run

# 清理输出目录
clean:
	@echo "Cleaning output..."
	@rm -rf output