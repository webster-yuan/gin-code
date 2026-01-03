package commands

import (
	"fmt"
	"gin/ds/Heap"
	"gin/ds/binary_tree"
	"gin/ds/lists"
	"gin/ds/slice"

	"github.com/spf13/cobra"
)

// DSCmd 定义ds子命令
var DSCmd = &cobra.Command{
	Use:   "ds",
	Short: "运行数据结构示例",
	Long:  `运行各种数据结构的示例代码，用于学习和测试。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("数据结构示例运行中...")
		fmt.Println("请选择具体的数据结构类型，例如：")
		fmt.Println("  - ds lists: 运行链表示例")
		fmt.Println("  - ds sorting: 运行排序算法示例")
	},
}

func init() {
	// 添加子命令
	DSCmd.AddCommand(listCmd)
	DSCmd.AddCommand(sortingCmd)
	DSCmd.AddCommand(sliceCmd)
	DSCmd.AddCommand(binaryTreeCmd)
	DSCmd.AddCommand(HeapCmd)
}

// listCmd 链表示例命令
var listCmd = &cobra.Command{
	Use:   "lists",
	Short: "运行链表示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行链表示例...")
		// 调用链表示例代码
		lists.ExampleLinkedList()
	},
}

// sliceCmd 顺序表示例命令
var sliceCmd = &cobra.Command{
	Use:   "slice",
	Short: "运行顺序表示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行顺序表示例...")
		// 调用链表示例代码
		slice.SeqListMain()
	},
}

// sortingCmd 排序算法示例命令
var sortingCmd = &cobra.Command{
	Use:   "sorting",
	Short: "运行排序算法示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行排序算法示例...")
		fmt.Println("排序算法示例功能开发中...")
	},
}

// binaryTreeCmd 二叉树示例命令
var binaryTreeCmd = &cobra.Command{
	Use:   "binary_tree",
	Short: "运行顺序表示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行顺序表示例...")
		// 调用链表示例代码
		binary_tree.TreeNodeMain()
	},
}

// HeapCmd 小根堆示例命令
var HeapCmd = &cobra.Command{
	Use:   "min_heap",
	Short: "运行小根堆示例",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("运行顺序表示例...")
		// 调用链表示例代码
		Heap.MinHeapMain()
	},
}
