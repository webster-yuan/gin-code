package advanced

import "fmt"

// Crane 起重机接口
type Crane interface {
	JackUp() string
	Hoist() string
}

type CraneA struct {
	work int
}

func (c CraneA) Work() {
	fmt.Println("CraneA is working")
}

func (c CraneA) JackUp() string {
	c.Work()
	return "Jack up"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "Hoist"
}

type CraneB struct {
	boot string
}

func (c CraneB) Boot() {
	fmt.Println("CranB is booting")
}

func (c CraneB) JackUp() string {
	c.Boot()
	return "Jack up"
}

func (c CraneB) Hoist() string {
	c.Boot()
	return "Hoist"
}

type ConstructionCompany struct {
	crane Crane
}

func (c ConstructionCompany) Builder() {
	fmt.Println(c.crane.JackUp())
	fmt.Println(c.crane.Hoist())
	fmt.Println("建筑完成")
}

// Any 接口内部没有方法集合，所以所有类型都是 Any接口的实现，因为所有类型都是空集的超集，所以 Any 接口可以保存任何类型的值
type Any interface {
}

func DoSomething(anything interface{}) interface{} {
	fmt.Println("DoSomething is called")
	return anything
}

// Equal 函数名可自定义（比如Equal），重点是[T comparable]这个约束
func Equal[T comparable](a, b T) bool {
	return a == b // 因为有comparable约束，所以能安全使用==
}

func EmptyInterface() {
	fmt.Println("EmptyInterface is called")
	// syg 等同于any
	type syg = interface{}
	var a syg
	var b syg
	a = 1
	b = "1"
	//在比较空接口时，会对其底层类型进行比较，如果类型不匹配的话则为false，其次才是值的比较
	fmt.Println(a == b)
	a = 1
	b = 1
	fmt.Println(a == b)

	fmt.Println(Equal("hello", "hello")) // 输出：true
}

func InterfaceMain() {
	// 使用起重机A
	company := ConstructionCompany{CraneA{1}}
	company.Builder()

	company.crane = CraneB{boot: "启动"}
	company.Builder()
	// 空接口
	var anything Any
	anything = 1
	println(anything)
	fmt.Printf("anything is %d\n", anything)
	anything = DoSomething(map[int]interface{}{1: "one", 2: "two"})
	fmt.Println(anything)
	EmptyInterface()
}

type Person interface {
	Say(string) string
	Walk(int)
}

// Man 拓展了自己的行为定义，所以要实现Man接口，必须同时实现Person中定义的所有方法
// 接口嵌入的本质：接口组合，优于类继承
type Man interface {
	Exercise()
	Person // 接口嵌入机制：会继承被嵌入接口的所有方法签名
}

// Webster 同时实现了 Man 和 Person 接口
type Webster struct {
	Name string
}

func (h Webster) Say(str string) string {
	return h.Name + " says " + str
}

func (h Webster) Walk(time int) {
	fmt.Printf("%s says %d times", h.Name, time)
}

func (h Webster) Exercise() {
	fmt.Println(h.Name, "is exercising")
}
func (h Webster) String() string {
	return h.Name
}

type Number int

func (n Number) Say(string) string {
	return fmt.Sprintf("%d", n)
}

func (n Number) Walk(int) {
	fmt.Printf("n is %d", n)
}
