package slice

import "fmt"

type SeqList struct {
	data []int
}

func New(capacity int) *SeqList {
	return &SeqList{make([]int, 0, capacity)}
}

func (sl *SeqList) Append(value int) {
	sl.data = append(sl.data, value)
}

// Insert 在index位置插入v，index∈[0,len]
func (sl *SeqList) Insert(position int, value int) {
	if position < 0 || position > len(sl.data) {
		panic("position out of range")
	}
	// 先扩容一个位置
	sl.data = append(sl.data, 0)
	// 向右移动位置
	copy(sl.data[position+1:], sl.data[position:])
	// 位置赋值
	sl.data[position] = value
}

// Delete 按照索引删除元素，返回被删除值
func (sl *SeqList) Delete(index int) int {
	value := sl.data[index]
	// 向左移动位置，同时把index位置的值进行了覆盖
	copy(sl.data[index:], sl.data[index+1:])
	// 将最后一个位置的空间去掉
	sl.data = sl.data[:len(sl.data)-1]
	return value
}

func String(sl *SeqList) string {
	return fmt.Sprintf("%v", sl.data)
}

func SeqListMain() {
	list := New(4)
	list.Append(10)
	list.Append(20)
	list.Insert(1, 15)
	fmt.Println("after insert:", list) // [10 15 20]
	list.Delete(0)
	fmt.Println("after delete:", list) // [15 20]
}
