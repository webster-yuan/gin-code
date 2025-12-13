总结一句话  
`strings.Replacer` 先一次性把多组“旧→新”规则编译成 Trie，然后**只扫描一遍输入**就能完成全部替换；  
而多次调用 `Replace`/`ReplaceAll` 每次都要重新扫一遍字符串，规则越多、重复次数越多，`Replacer` 越快。

---------------- 代码对比 ----------------
```go
// 要替换 3 组词
old := "foo bar baz foo bar baz"

// 方法1：Replace 逐次做 → 扫 3 遍
s1 := strings.Replace(old, "foo", "FOO", -1)
s1 = strings.Replace(s1, "bar", "BAR", -1)
s1 = strings.Replace(s1, "baz", "BAZ", -1)

// 方法2：Replacer 一次做 → 扫 1 遍
rep := strings.NewReplacer(
    "foo", "FOO",
    "bar", "BAR",
    "baz", "BAZ",
)
s2 := rep.Replace(old)

fmt.Println(s1 == s2) // true
```

同样结果，但
- `Replace` 链式调用：每步生成一个新字符串，共扫描 3 次。
- `Replacer`：预处理一次，运行时只扫描 1 次，且可复用 `rep` 无数次。