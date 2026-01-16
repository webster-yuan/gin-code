#!/bin/bash

# 测试 API 的脚本

echo "=== 测试用户 API ==="
echo ""

# 1. 创建用户
echo "1. 创建用户..."
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }'
echo ""
echo ""

# 2. 获取所有用户
echo "2. 获取所有用户..."
curl -X GET http://localhost:8080/api/v1/users
echo ""
echo ""

# 3. 获取单个用户（假设ID为1）
echo "3. 获取用户ID=1..."
curl -X GET http://localhost:8080/api/v1/users/1
echo ""
echo ""

# 4. 更新用户
echo "4. 更新用户ID=1..."
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三（已更新）",
    "age": 26
  }'
echo ""
echo ""

# 5. 再次获取用户验证更新
echo "5. 验证更新后的用户..."
curl -X GET http://localhost:8080/api/v1/users/1
echo ""
echo ""

echo "=== 测试完成 ==="
