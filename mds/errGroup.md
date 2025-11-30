`errgroup` 的名字直接暴露了它的**设计目的**：

- **group** → 把**一组 goroutine** 聚到一起管理；
- **err** → 只要**任意一个 goroutine 返回错误**，整个组就**立刻被取消**并把该错误**回传**给调用者。

---

### 一句话职责
> “并发跑一堆任务，**第一个出错就停全员**，并把错误**集中**交给我。”

---

### 对比原生 `sync.WaitGroup`
| 特性 | sync.WaitGroup | golang.org/x/sync/errgroup |
|---|---|---|
| 只等结束 | ✅ | ✅ |
| 能取消 sibling goroutine | ❌ | ✅（`context.Cancel`） |
| 能把第一个错误带回来 | ❌（自己收集） | ✅（`g.Wait()` 直接返回） |

---

### 命名类比（一看就懂）
- `errgroup` = **带错误处理的 WaitGroup**
- `atomic` = **原子操作包**
- `httptest` = **HTTP 测试包**

Go 官方习惯用**“功能关键词”**拼包名，**见名知义**，减少文档负担。

---

### 一句话
叫 `errgroup` 就是告诉你：  
“这玩意儿**专门管一组 goroutine 的错误**，谁崩就停谁，最后把错误**一起**带回来。”