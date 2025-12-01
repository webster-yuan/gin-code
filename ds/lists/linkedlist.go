package lists

import "fmt"

// Node 链表节点结构
type Node struct {
	Value int
	Next  *Node
}

// LinkedList 链表结构
type LinkedList struct {
	Head *Node
}

// NewLinkedList 创建新链表
func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

// Append 添加节点到链表末尾
func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Print 打印链表内容
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// ExampleLinkedList 运行链表示例
func ExampleLinkedList() {
	ll := NewLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	fmt.Println("链表内容:")
	ll.Print()
}
