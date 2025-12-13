`strings.Builder` 就是 Go 标准库给你提供的**“可变形字符串缓冲区”**，作用跟 C++ 里的

```cpp
std::ostringstream oss;
oss << "hello";
std::string s = oss.str();
```

几乎一模一样，但实现得更轻量、零拷贝：

1. 底层是一个 `[]byte` 切片，容量不够时自动 2× 扩容。
2. 所有 `Write*` 方法都直接往这片内存里追加，**不会**像普通 `str = str + "xxx"` 那样产生中间字符串。
3. 写完以后调用 `.String()` 一次性把底层 `[]byte` 头变成 `string` 返回，全程只拷贝一次。

因此它专用于**高频拼接**场景，比 `+` 或 `fmt.Sprintf` 快且省内存。

简单记忆：  
C++ 的 `std::ostringstream` / Java 的 `StringBuilder` / Rust 的 `String`  
→ 在 Go 里就叫 `strings.Builder`，用法也是  
“创建 → 写 → `.String()` 取结果”，没有魔法。