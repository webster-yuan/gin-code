
---

# 一、Viper（配置管理）——工程必备函数清单

> 目标：**能在真实项目中稳定用配置，不踩坑**

## ✅ 必须掌握（90% 项目只用这些）

### 1️⃣ `viper.SetConfigName`

```go
viper.SetConfigName("config")
```

**干什么的**
指定配置文件名（不含后缀）

**什么时候用**
👉 项目用配置文件时 *必用*

---

### 2️⃣ `viper.SetConfigType`

```go
viper.SetConfigType("yaml")
```

**干什么的**
指定配置文件格式

**什么时候用**

* 配置文件没有 `.yaml` 后缀时必用
* 否则可以省略

---

### 3️⃣ `viper.AddConfigPath`

```go
viper.AddConfigPath("./")
viper.AddConfigPath("/etc/xxx/")
```

**干什么的**
告诉 Viper 去哪些目录找配置文件

**工程惯例**

* `./` → 本地开发
* `/etc/项目名/` → 服务器

---

### 4️⃣ `viper.AutomaticEnv`

```go
viper.AutomaticEnv()
```

**干什么的**
让配置支持 **环境变量覆盖**

**工程意义（非常重要）**

* 容器化
* CI/CD
* k8s

👉 **所有工程项目必开**

---

### 5️⃣ `viper.SetDefault`

```go
viper.SetDefault("server.port", 8080)
```

**干什么的**
配置兜底值（防止没配置程序直接挂）

**工程建议**

* 只给「安全默认值」
* 不要给敏感信息（如密码）

---

### 6️⃣ `viper.ReadInConfig`

```go
err := viper.ReadInConfig()
```

**干什么的**
真正读取配置文件

**工程习惯**

```go
if err != nil {
	log.Println("配置文件不存在，使用默认/环境变量")
}
```

👉 **通常不 Fatal**

---

### 7️⃣ `viper.Unmarshal`

```go
viper.Unmarshal(&cfg)
```

**干什么的**
把配置映射成结构体

**工程规则**

* 失败 = 程序配置错误
* 直接 `Fatal`

---

## ❌ 可以忽略（现阶段完全不需要）

* `WatchConfig`
* `OnConfigChange`
* `SetEnvKeyReplacer`
* `BindEnv`
* 多配置文件 merge

> 等你做到「配置热更新」或「超复杂环境」再看

---

