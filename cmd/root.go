package cmd

import (
	"gin/cmd/commands"
	"log"

	"github.com/spf13/cobra"
)

// rootCmd 表示没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "gin",
	Short: "Gin API服务与Go语法学习工具",
	Long: `一个集成了Gin框架API服务和Go语言学习示例的工具。
	
使用子命令来区分不同功能：
  - server: 运行Gin API服务
  - ds: 运行数据结构示例
  - examples: 运行Go语法示例`,
}

// Execute 执行根命令
func Execute() {
	// 添加子命令
	rootCmd.AddCommand(commands.ServerCmd)
	rootCmd.AddCommand(commands.DSCmd)
	rootCmd.AddCommand(commands.ExamplesCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
