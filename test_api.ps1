# PowerShell 测试脚本

Write-Host "=== 测试用户 API ===" -ForegroundColor Green
Write-Host ""

# 1. 创建用户
Write-Host "1. 创建用户..." -ForegroundColor Yellow
$body = @{
    name = "张三"
    email = "zhangsan@example.com"
    age = 25
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users" -Method POST -Body $body -ContentType "application/json"
Write-Host ""

# 2. 获取所有用户
Write-Host "2. 获取所有用户..." -ForegroundColor Yellow
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users" -Method GET
Write-Host ""

# 3. 获取单个用户（假设ID为1）
Write-Host "3. 获取用户ID=1..." -ForegroundColor Yellow
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users/1" -Method GET
Write-Host ""

# 4. 更新用户
Write-Host "4. 更新用户ID=1..." -ForegroundColor Yellow
$updateBody = @{
    name = "张三（已更新）"
    age = 26
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users/1" -Method PUT -Body $updateBody -ContentType "application/json"
Write-Host ""

# 5. 再次获取用户验证更新
Write-Host "5. 验证更新后的用户..." -ForegroundColor Yellow
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/users/1" -Method GET
Write-Host ""

Write-Host "=== 测试完成 ===" -ForegroundColor Green
