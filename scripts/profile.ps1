# PowerShell 性能分析脚本

Write-Host "=== Go 性能分析工具 ===" -ForegroundColor Green
Write-Host ""
Write-Host "请选择分析类型："
Write-Host "1. CPU 性能分析（30秒采样）"
Write-Host "2. 内存性能分析"
Write-Host "3. Goroutine 分析"
Write-Host "4. 阻塞分析"
Write-Host "5. 互斥锁分析"
Write-Host "6. 启动 Web 界面（CPU，30秒）"
Write-Host "7. 启动 Web 界面（内存）"
Write-Host ""

$choice = Read-Host "请输入选项 (1-7)"

switch ($choice) {
    "1" {
        Write-Host "正在生成 CPU 性能分析（30秒）..." -ForegroundColor Yellow
        go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
    }
    "2" {
        Write-Host "正在生成内存性能分析..." -ForegroundColor Yellow
        go tool pprof http://localhost:6060/debug/pprof/heap
    }
    "3" {
        Write-Host "正在生成 Goroutine 分析..." -ForegroundColor Yellow
        go tool pprof http://localhost:6060/debug/pprof/goroutine
    }
    "4" {
        Write-Host "正在生成阻塞分析..." -ForegroundColor Yellow
        go tool pprof http://localhost:6060/debug/pprof/block
    }
    "5" {
        Write-Host "正在生成互斥锁分析..." -ForegroundColor Yellow
        go tool pprof http://localhost:6060/debug/pprof/mutex
    }
    "6" {
        Write-Host "正在启动 Web 界面（CPU，30秒采样）..." -ForegroundColor Yellow
        Write-Host "请在浏览器中访问: http://localhost:8081" -ForegroundColor Cyan
        go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
    }
    "7" {
        Write-Host "正在启动 Web 界面（内存）..." -ForegroundColor Yellow
        Write-Host "请在浏览器中访问: http://localhost:8081" -ForegroundColor Cyan
        go tool pprof -http=:8081 http://localhost:6060/debug/pprof/heap
    }
    default {
        Write-Host "无效选项" -ForegroundColor Red
        exit 1
    }
}
