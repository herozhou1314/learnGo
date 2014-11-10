package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("hello`")

	vertex := map[string]Vertex{"123": Vertex{123, 123}}

	fmt.Println(vertex)
	fmt.Println(vertex["123"])
	m := map[string]string{}
	m["key1"] = "hello"
	m["key2"] = "hello"
	fmt.Println(m)
	delete(m, "key1")
	fmt.Println(m)
	maps := map[string]VertexS{"key": VertexS{"hello", "world"}}
	maps["key1"] = VertexS{"123", "456"}
	fmt.Println("hello", maps, maps["key1"])
	total := func(x, y int) int {
		fmt.Println("add")
		return x + y
	}
	fmt.Println(total(1, 2))
	fmt.Println(adder(123))

	m1 := Vertex{1, 2}
	fmt.Println(m1.addBy())
	m1.pointer()
	fmt.Println(m1)
}

type VertexS struct {
	m, n string
}

func adder(n int) func() {
	sum := n
	a := func() { //把匿名函数作为值赋给变量a (Go 不允许函数嵌套。
		//然而你可以利用匿名函数实现函数嵌套)
		fmt.Println(sum + 1) //调用本函数外的变量
	} //这里没有()匿名函数不会马上执行
	return a
}


/****
在结构体类型上定义方法
*/
func (v *Vertex) addBy() float64 {
	return v.Lat + v.Long
}

/****
在结构体类型上定义方法
*/
func (v *Vertex) minus() float64 {
	return v.Lat - v.Long
}
/****
在结构体类型上定义方法
*/
func (v *Vertex) pointer()  {
	 v.Lat = 1000
	 v.Long = 1001
}
