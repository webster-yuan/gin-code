package stack

type Stack []int

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	// 模拟先进后出，所以返回末尾的元素
	value := (*s)[len(*s)-1]
	// 将去掉元素的最后位置去掉
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
