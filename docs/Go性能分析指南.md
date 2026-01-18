# Go 性能分析指南

Go 语言提供了强大的性能分析工具，可以分析程序的 CPU 使用、内存分配、阻塞情况等。

## 性能分析类型

Go 支持以下几种性能分析：

1. **CPU 性能分析**（CPU Profiling）- 分析函数执行时间
2. **内存性能分析**（Memory Profiling）- 分析内存分配
3. **阻塞性能分析**（Block Profiling）- 分析 goroutine 阻塞
4. **互斥锁性能分析**（Mutex Profiling）- 分析锁竞争
5. **Goroutine 性能分析**（Goroutine Profiling）- 分析 goroutine 状态

## 方法一：使用 `go test` 生成性能分析数据

### CPU 性能分析

```bash
# 生成 CPU 性能分析文件
go test -cpuprofile=cpu.prof ./internal/service/...

# 查看 CPU 性能分析
go tool pprof cpu.prof
```

### 内存性能分析

```bash
# 生成内存性能分析文件
go test -memprofile=mem.prof ./internal/service/...

# 查看内存性能分析
go tool pprof mem.prof
```

### 同时生成多种分析

```bash
# 生成 CPU 和内存分析
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./internal/service/...

# 生成所有类型的分析
go test \
  -cpuprofile=cpu.prof \
  -memprofile=mem.prof \
  -blockprofile=block.prof \
  -mutexprofile=mutex.prof \
  ./internal/service/...
```

## 方法二：使用 `net/http/pprof`（推荐）

这是最常用的方法，可以在运行时通过 HTTP 端点实时分析。

### 1. 在代码中启用 pprof

```go
import _ "net/http/pprof"

func main() {
    // 启动 pprof HTTP 服务器（默认在 :6060）
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // 你的应用代码...
}
```

### 2. 访问性能分析端点

启动应用后，访问以下 URL：

- **CPU 性能分析**：`http://localhost:6060/debug/pprof/profile?seconds=30`
  - `seconds` 参数指定采样时间（秒）
  - 会自动下载 `profile` 文件

- **内存性能分析**：`http://localhost:6060/debug/pprof/heap`
  - 下载 `heap` 文件

- **Goroutine 分析**：`http://localhost:6060/debug/pprof/goroutine`
  - 查看所有 goroutine 的状态

- **阻塞分析**：`http://localhost:6060/debug/pprof/block`

- **互斥锁分析**：`http://localhost:6060/debug/pprof/mutex`

- **查看所有可用端点**：`http://localhost:6060/debug/pprof/`

### 3. 使用命令行工具分析

```bash
# CPU 性能分析（采样 30 秒）
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# 内存性能分析
go tool pprof http://localhost:6060/debug/pprof/heap

# 实时查看（不下载文件）
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```

## 方法三：使用 `runtime/pprof` 手动生成

在代码中手动生成性能分析数据：

```go
import (
    "os"
    "runtime/pprof"
)

func main() {
    // CPU 性能分析
    f, _ := os.Create("cpu.prof")
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // 你的代码...
    
    // 内存性能分析
    mf, _ := os.Create("mem.prof")
    defer mf.Close()
    pprof.WriteHeapProfile(mf)
}
```

## 查看和分析性能数据

### 1. 交互式命令行查看

```bash
# 进入交互模式
go tool pprof cpu.prof

# 常用命令：
# top - 显示占用最多的函数
# top10 - 显示前 10 个
# list 函数名 - 查看函数详细代码
# web - 生成调用图（需要安装 graphviz）
# exit - 退出
```

### 2. Web 界面查看（推荐）

```bash
# 启动 Web 界面（默认 :8080）
go tool pprof -http=:8081 cpu.prof

# 或者直接分析远程端点
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```

然后在浏览器访问 `http://localhost:8081`，可以看到：
- **Top** - 函数调用统计
- **Graph** - 调用关系图
- **Flame Graph** - 火焰图
- **Source** - 源代码视图

### 3. 生成报告

```bash
# 生成文本报告
go tool pprof -text cpu.prof > cpu.txt

# 生成 PDF 报告（需要 graphviz）
go tool pprof -pdf cpu.prof > cpu.pdf

# 生成 SVG 图
go tool pprof -svg cpu.prof > cpu.svg
```

## 实际示例

### 示例 1：分析测试性能

```bash
# 运行测试并生成 CPU 分析
go test -cpuprofile=cpu.prof -bench=. ./internal/service/...

# 查看分析结果
go tool pprof cpu.prof

# 在交互模式中输入：
# top10        # 查看前 10 个最耗时的函数
# list CreateUser  # 查看 CreateUser 函数的详细分析
```

### 示例 2：分析运行中的程序

```bash
# 1. 启动应用（已启用 pprof）
go run main.go server

# 2. 在另一个终端生成 30 秒的 CPU 分析
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# 3. 或者使用 Web 界面
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```

### 示例 3：对比分析

```bash
# 生成第一个分析
go test -cpuprofile=cpu1.prof ./...

# 修改代码后生成第二个分析
go test -cpuprofile=cpu2.prof ./...

# 对比分析
go tool pprof -base=cpu1.prof cpu2.prof
```

## 常用命令速查

### 生成分析数据

```bash
# 测试时生成
go test -cpuprofile=cpu.prof -memprofile=mem.prof ./...

# 基准测试时生成
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof ./...

# 分析运行中的程序
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

### 查看分析数据

```bash
# 命令行交互模式
go tool pprof cpu.prof

# Web 界面
go tool pprof -http=:8081 cpu.prof

# 文本报告
go tool pprof -text cpu.prof

# 生成调用图（需要 graphviz）
go tool pprof -web cpu.prof
```

### 在 pprof 交互模式中的命令

```
top          # 显示占用最多的函数
top10        # 显示前 10 个
list 函数名   # 查看函数详细代码和耗时
web          # 生成调用图（需要 graphviz）
png          # 生成 PNG 图
svg          # 生成 SVG 图
help         # 显示帮助
exit         # 退出
```

## 性能分析输出解读

### CPU 性能分析输出示例

```
(pprof) top10
Showing nodes accounting for 200ms, 50.00% of 400ms total
      flat  flat%   sum%        cum   cum%
      80ms  20.00%  20.00%     120ms  30.00%  runtime.mallocgc
      60ms  15.00%  35.00%      60ms  15.00%  encoding/json.Marshal
      40ms  10.00%  45.00%     200ms  50.00%  main.handler
      ...
```

**字段说明：**
- `flat` - 函数自身执行时间（不包括调用其他函数）
- `flat%` - flat 占总时间的百分比
- `cum` - 累计时间（包括调用其他函数）
- `cum%` - cum 占总时间的百分比

### 内存性能分析输出示例

```
(pprof) top10
Showing nodes accounting for 50MB, 50.00% of 100MB total
      flat  flat%   sum%        cum   cum%
    20MB  20.00%  20.00%     30MB  30.00%  runtime.mallocgc
    15MB  15.00%  35.00%     15MB  15.00%  encoding/json.Marshal
    10MB  10.00%  45.00%     50MB  50.00%  main.handler
      ...
```

## 性能优化建议

1. **识别热点函数**：使用 `top` 命令找出最耗时的函数
2. **分析调用链**：使用 `web` 或 `list` 查看完整的调用关系
3. **内存优化**：使用内存分析找出内存分配热点
4. **并发优化**：使用 goroutine 分析找出阻塞点

## 安装 graphviz（可选）

如果需要生成调用图，需要安装 graphviz：

```bash
# Windows (使用 Chocolatey)
choco install graphviz

# macOS
brew install graphviz

# Ubuntu/Debian
sudo apt-get install graphviz
```

## 注意事项

1. **性能开销**：性能分析会带来一定的性能开销（通常 5-10%）
2. **采样时间**：CPU 分析需要足够的采样时间才能准确
3. **生产环境**：谨慎在生产环境启用 pprof，建议使用独立的监控端口
4. **文件大小**：性能分析文件可能很大，注意磁盘空间

## 参考资料

- [Go 官方文档 - Profiling](https://golang.org/doc/diagnostics#profiling)
- [pprof 工具文档](https://github.com/google/pprof)
- [Go 性能优化实战](https://github.com/golang/go/wiki/Performance)
