package queue

type Queue []int

// Enqueue 入队
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value)
}

// Dequeue 出队
func (q *Queue) Dequeue() int {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

// IsEmpty 是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
