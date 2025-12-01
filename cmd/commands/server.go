package commands

import (
	"gin/cmd/server"

	"github.com/spf13/cobra"
)

// ServerCmd 定义server子命令
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "运行Gin API服务",
	Long:  `启动Gin框架的API服务，包括主服务和多个子服务。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 调用现有的server.Main()函数
		server.Main()
	},
}
