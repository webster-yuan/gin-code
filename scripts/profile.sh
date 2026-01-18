#!/bin/bash

# Go 性能分析脚本

echo "=== Go 性能分析工具 ==="
echo ""
echo "请选择分析类型："
echo "1. CPU 性能分析（30秒采样）"
echo "2. 内存性能分析"
echo "3. Goroutine 分析"
echo "4. 阻塞分析"
echo "5. 互斥锁分析"
echo "6. 启动 Web 界面（CPU，30秒）"
echo "7. 启动 Web 界面（内存）"
echo ""

read -p "请输入选项 (1-7): " choice

case $choice in
    1)
        echo "正在生成 CPU 性能分析（30秒）..."
        go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
        ;;
    2)
        echo "正在生成内存性能分析..."
        go tool pprof http://localhost:6060/debug/pprof/heap
        ;;
    3)
        echo "正在生成 Goroutine 分析..."
        go tool pprof http://localhost:6060/debug/pprof/goroutine
        ;;
    4)
        echo "正在生成阻塞分析..."
        go tool pprof http://localhost:6060/debug/pprof/block
        ;;
    5)
        echo "正在生成互斥锁分析..."
        go tool pprof http://localhost:6060/debug/pprof/mutex
        ;;
    6)
        echo "正在启动 Web 界面（CPU，30秒采样）..."
        echo "请在浏览器中访问: http://localhost:8081"
        go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
        ;;
    7)
        echo "正在启动 Web 界面（内存）..."
        echo "请在浏览器中访问: http://localhost:8081"
        go tool pprof -http=:8081 http://localhost:6060/debug/pprof/heap
        ;;
    *)
        echo "无效选项"
        exit 1
        ;;
esac
