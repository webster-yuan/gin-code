package custom_map

import "fmt"

// Go 语言内置map，平均O(1)时间读取

func MapMain() {
	m := make(map[string]int)
	m["apple"] = 34
	m["banana"] = 2
	cnt, ok := m["apple"]
	if ok == false {
		panic("apple")
	}
	fmt.Printf("%d\n", cnt)
	delete(m, "apple")
	fmt.Printf("%#v\n", m)
}
