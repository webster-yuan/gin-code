package commands

import (
	"fmt"
	"gin/examples/basics"

	"github.com/spf13/cobra"
)

// ExamplesCmd 定义examples子命令
var ExamplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "运行Go语法示例",
	Long:  `运行各种Go语言语法特性的示例代码，用于学习和测试。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go语法示例运行中...")
		fmt.Println("请选择具体的示例类型，例如：")
		fmt.Println("  - examples basics: 运行基础语法示例")
		fmt.Println("  - examples concurrency: 运行并发示例")
	},
}

func init() {
	// 添加子命令
	ExamplesCmd.AddCommand(basicsCmd)
	ExamplesCmd.AddCommand(concurrencyCmd)
}

// basicsCmd 基础语法示例命令
var basicsCmd = &cobra.Command{
	Use:   "basics",
	Short: "运行基础语法示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行基础语法示例...")
		// 调用结构体示例代码
		basics.ExampleStructs()
	},
}

// concurrencyCmd 并发示例命令
var concurrencyCmd = &cobra.Command{
	Use:   "concurrency",
	Short: "运行并发示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行并发示例...")
		fmt.Println("并发示例功能开发中...")
	},
}
