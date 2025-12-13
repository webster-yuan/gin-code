## 项目命令结构概览
项目使用cobra库实现了以下命令结构：
- 根命令：`gin`
- 子命令：
  - `server`: 运行Gin API服务
  - `ds`: 运行数据结构示例
    - `lists`: 运行链表示例
    - `sorting`: 运行排序算法示例
  - `examples`: 运行Go语法示例
    - `basics`: 运行基础语法示例
    - `concurrency`: 运行并发示例

## Goland 2025配置运行图标的步骤

### 1. 配置主入口文件
首先，我们需要为项目的主入口创建一个运行配置：

1. 点击顶部菜单栏的 **Run** -> **Edit Configurations...**
2. 点击左上角的 **+** 按钮，选择 **Go Build**
3. 配置基本信息：
   - **Name**: 可以设置为 `gin-main`
   - **Run kind**: 选择 **`File`**
   - **File**: 选择项目根目录下的 `E:\gitbub-code\gin\main.go` 文件
   - **Working directory**: 设置为您的项目根目录 `E:\gitbub-code\gin`
   - **Output directory**: `E:\gitbub-code\gin\bin`

4. 点击 **Apply** 保存配置

### 2. 配置server子命令
为了运行API服务部分：

1. 再次点击 **+** 按钮，选择 **Go Build**
2. 配置信息：
   - **Name**: 设置为 `gin-server`
   - **Run kind**: 选择 **`File`**
   - **Files**: 输入 `E:\gitbub-code\gin\main.go`
   - **Working directory**: `E:\gitbub-code\gin`
   - **Program arguments**: 输入 `server`
   - **Output directory**: `E:\gitbub-code\gin\bin`
3. 点击 **Apply** 保存配置

### 3. 配置ds相关子命令
为数据结构示例创建配置：

#### 配置ds根命令：
1. 点击 **+** 按钮，选择 **Go Build**
2. 配置：
   - **Name**: `gin-ds`
   - **Run kind**: `File`
   - **Files**: 输入 `E:\gitbub-code\gin\main.go`
   - **Working directory**: `E:\gitbub-code\gin`
   - **Program arguments**: `ds`
   - **Output directory**: `E:\gitbub-code\gin\bin`

#### 配置链表示例：
1. 点击 **+** 按钮，选择 **Go Build**
2. 配置：
   - **Name**: `gin-ds-lists`
   - **Run kind**: `File`
   - **Files**: 输入 `E:\gitbub-code\gin\main.go`
   - **Working directory**: `E:\gitbub-code\gin`
   - **Program arguments**: `ds lists`
   - **Output directory**: `E:\gitbub-code\gin\bin`

### 4. 配置examples相关子命令
为Go语法示例创建配置：

#### 配置基础语法示例：
1. 点击 **+** 按钮，选择 **Go Build**
2. 配置：
   - **Name**: `gin-examples-basics`
   - **Run kind**: `File`
   - **Files**: 输入 `E:\gitbub-code\gin\main.go`
   - **Working directory**: `E:\gitbub-code\gin`
   - **Program arguments**: `examples basics`
   - **Output directory**: `E:\gitbub-code\gin\bin`

### 5. 使用配置运行
配置完成后，您可以通过以下方式使用：

1. 在顶部工具栏的运行配置下拉菜单中选择相应的配置
2. 点击绿色的运行按钮或使用快捷键（通常是Shift+F10）运行
3. 您也可以在项目视图中右键点击相应的配置并选择运行

### 6. 为配置添加图标（可选）
Goland 2025支持为运行配置添加自定义图标：

1. 在配置编辑窗口中，点击配置名称旁边的图标
2. 在弹出的对话框中选择一个预设图标或上传自定义图标
3. 点击 **OK** 保存

这样配置后，就可以通过点击图标轻松运行项目的不同部分，而不需要每次都打开命令行并输入命令了。如果您需要添加更多子命令的配置，只需按照类似的步骤操作即可。
        