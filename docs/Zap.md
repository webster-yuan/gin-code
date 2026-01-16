# 二、Zap（日志）——工程必备函数清单

> 目标：**日志统一、可控、线上可用**

---

## ✅ 必须掌握（核心 10%）

### 1️⃣ `zap.NewProduction / NewDevelopment`

```go
logger, _ := zap.NewProduction()
```

```go
logger, _ := zap.NewDevelopment()
```

**干什么的**

* 创建日志实例

**工程惯例**

* 开发环境：`NewDevelopment`
* 生产环境：`NewProduction`

---

### 2️⃣ `logger.Info / Error / Warn`

```go
logger.Info("server started")
logger.Error("db error", zap.Error(err))
```

**干什么的**

* 写日志

**工程规则**

* Info：业务流程
* Warn：可恢复异常
* Error：明确错误

---

### 3️⃣ `zap.Error / zap.String / zap.Int`

```go
zap.Error(err)
zap.String("user_id", uid)
```

**干什么的**

* 结构化字段（JSON 日志）

**工程意义**
👉 日志不是给人看的，是给 **ELK / Loki / Datadog** 看的

---

### 4️⃣ `logger.Sync()`

```go
defer logger.Sync()
```

**干什么的**

* 刷新日志缓冲

**工程规则**

* main 函数里一定要有

---

### 5️⃣ `zap.ReplaceGlobals`

```go
zap.ReplaceGlobals(logger)
```

**干什么的**

* 把 logger 设为全局

之后可以直接用：

```go
zap.L().Info("xxx")
zap.S().Errorf("xxx")
```

**工程意义**
👉 避免 logger 到处传

---

## 🧠 工程里 zap 的“标准姿势”

```go
func InitLogger(cfg *Config) {
	var logger *zap.Logger

	if cfg.Env == "prod" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	zap.ReplaceGlobals(logger)
}
```

然后全项目：

```go
zap.L().Info("server start")
```

---

## ❌ 现阶段不用学的 Zap 内容

* Core / Encoder / Syncer
* Lumberjack（日志切割）
* 自定义编码器
* Hook

> 这些是 **日志平台 / 中台 / SRE 级别内容**

---

# 三、Gin + Viper + Zap（工程组合记忆法）

只记一句话就够了：

> **Viper 管“配置怎么来”
> Zap 管“日志怎么打”
> Gin 管“请求怎么走”**

### 启动顺序（固定套路）

```go
func main() {
	cfg := LoadConfig()     // viper
	InitLogger(cfg)        // zap
	r := InitRouter(cfg)   // gin
	r.Run()
}
```

---
