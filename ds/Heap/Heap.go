package Heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

// Less 实现Less方法，告诉标准库我们实现的是小根堆 Min-Heap,根永远是最小值
func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop 绝不直接返回最小值；它只是“删末尾”的纯工具
// 最小值在标准库内部先被保存，然后由标准库完成“搬末尾→下沉→返回最小值”的完整闭环
// 正因为标准库需要“拿到末尾元素”再“下沉”，所以才要求你实现 Pop 接口——它只信任你来做“从你自己的数据结构里删掉末尾并交给我”这一步
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MinHeapMain() {
	h := &IntHeap{5, 3, 4, 7, 6, 9}
	heap.Init(h)    // ① 把切片堆化，O(n)
	heap.Push(h, 1) // ② 插入 1，自动上浮到根，O(log n)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // ③ 依次弹出最小值
	}
}
