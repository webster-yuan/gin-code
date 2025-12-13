package advanced

import (
	"fmt"
)

func Sum[T int | float64](a, b T) T {
	return a + b
}

// GeneticsSlice 泛型切片，限定元素类型为int、float64、string
type GeneticsSlice[T int | float64 | string] []T

// SumSlice1 无泛型参数，直接限定参数为GeneticsSlice[int]
func SumSlice1(slice GeneticsSlice[int]) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumSlice2 有泛型参数，适配类型int、float64、string
func SumSlice2[T int | float64 | string](slice GeneticsSlice[T]) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// GenericsMap 泛型映射，限定键为可比较类型，值为int、float64、string
type GenericsMap[K comparable, V int | float64 | string] map[K]V

// GenericsStruct 泛型结构体，限定字段Id为int或string
type GenericsStruct[T int | string] struct {
	Name string
	Id   T
}

// Company 泛型结构体，限定字段Id为int或string，Stuff为泛型切片
type Company[T int | string, S []T] struct {
	Name  string
	Id    T
	Stuff S
}

// SayAble 泛型接口，限定方法Say返回值为int或string
type SayAble[T string | int] interface {
	Say(msg T) T
}

type Yuan[T string | int] struct {
	msg T
}

func (y Yuan[T]) Say(msg T) T {
	fmt.Println(msg)
	y.msg = msg
	return y.msg
}

// Greeter 非泛型接口，来用做参数限制，只有实现了这个接口的类型，才能作为参数传入
type Greeter interface {
	Hello() string
}

type Dog struct{}

func (d Dog) Hello() string {
	return "Hello Dog"
}

type Cat struct{}

func (c Cat) Hello() string {
	return "Hello Cat"
}

func Greet[T Greeter](obg T) {
	fmt.Println(obg.Hello())
}

// Assert 泛型断言，取代v.(int)这种断言容易出现的panic
func Assert[T any](v any) (bool, T) {
	var av T // 先给个T类型的零值，万一失败就用它
	if v == nil {
		return false, av
	}
	av, ok := v.(T)
	return ok, av
}

// 类型集

// SingedInt 有符号整数
type SingedInt interface {
	int8 | int16 | int32 | int64
}

// UnSignedInt 无符号整数
type UnSignedInt interface {
	uint8 | uint16 | uint32 | uint64
}

// Integer 整数
type Integer interface {
	SingedInt | UnSignedInt
}

// Intersection 交集，限定类型为Integer和SingedInt的交集
type Intersection interface {
	Integer
	SingedInt
}

// EmptySet 空集，因为 SingedInt UnSignedInt 之间交集为空
type EmptySet interface {
	SingedInt
	UnSignedInt
}

func DoIntersection[T Intersection](n T) T {
	return n
}

func DoEmptySet[T EmptySet](n T) T {
	return n
}

// Do 空接口是所有类型集的集合，不过一般都用 any ，代替空接口 interface{}
func Do[T interface{}](n T) T {
	return n
}

//type Int interface {
//	int8 | int16 | int32 | int64
//}

type Int interface {
	~int8 | ~int16 | ~int32 | ~int64
}

// TinyInt 当使用 type 关键字声明了一个新的类型时，即便其底层类型包含在类型集内，当传入时也依旧会无法通过编译。
//
//	“~” 是底层通行证，去掉它就只能收身份证（int8），不收改名卡（TinyInt）。
type TinyInt int8

func DoTinyInt[T Int](n T) T {
	return n
}

func GenericsMain() {
	DoTinyInt[TinyInt](1)

	//Do[struct{}](struct{}{})
	//
	//// DoEmptySet[int8](2) //Cannot use int8 as the type EmptySet Type does not implement constraint EmptySet because constraint type set is empty
	//
	//DoIntersection[int8](2)
	//Do[uint8](2) //Type does not implement constraint Intersection because type is not included in type set (int8, int16, int32, int64)

	//var x any = 42
	//if ok, _ := Assert[int](x); ok {
	//	fmt.Printf("x 是 int 类型的值")
	//} else {
	//	fmt.Printf("x 不是 int 类型的值")
	//}

	//dog := Dog{}
	//cat := Cat{}
	//Greet(dog)
	//Greet(cat)

	//var s SayAble[string]
	//s = Yuan[string]{"hello"}
	//ret := s.Say("yuan")
	//fmt.Printf("type is %T, value is %v\n", ret, ret)
	//
	//ret1 := Sum[int](1, 2)
	//fmt.Printf("type is %T, value is %v\n", ret1, ret1)
	//ret2 := Sum(1.2, 1)
	//fmt.Printf("type is %T, value is %v\n", ret2, ret2)
	//var geneticsSlice GeneticsSlice[int]
	//geneticsSlice = GeneticsSlice[int]{1, 2, 3}
	//fmt.Printf("type is %T, value is %v\n", geneticsSlice, geneticsSlice)
	//genericsMap1 := GenericsMap[int, string]{1: "1", 2: "2"}
	//fmt.Printf("type is %T, value is %v\n", genericsMap1, genericsMap1)
	//genericsMap2 := make(GenericsMap[int, string], 2)
	//genericsMap2[4] = "1"
	//genericsMap2[1] = "2"
	//fmt.Printf("type is %T, value is %v\n", genericsMap2, genericsMap2)
	//genericsStruct1 := GenericsStruct[int]{Name: "1", Id: 1}
	//fmt.Printf("type is %T, value is %v\n", genericsStruct1, genericsStruct1)
}
