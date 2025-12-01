package basics

import "fmt"

// Person 结构体示例
type Person struct {
	Name string
	Age  int
}

// ExampleStructs 运行结构体示例
func ExampleStructs() {
	// 基本结构体初始化
	p1 := Person{Name: "张三", Age: 25}
	fmt.Printf("p1: %+v\n", p1)

	// 匿名结构体
	p2 := struct {
		Name  string
		Email string
	}{
		Name:  "李四",
		Email: "lisi@example.com",
	}
	fmt.Printf("p2: %+v\n", p2)

	// 结构体指针
	p3 := &Person{Name: "王五", Age: 30}
	fmt.Printf("p3: %+v\n", p3)
}
