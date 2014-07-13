package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

/**给stuct 定义方法*/
func (v *Vertex) abs() float64 {
	return v.X + v.Y
}
type mystring string
func (f mystring) Abs() mystring {

	return f
}

func main() {
	fmt.Println("hello--->>")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, ok := range pow {
		fmt.Println(index, ok)
	}
	m := make(map[string]Vertex)
	m["baijing"] = Vertex{40.68433, -74.39967}
	m["jingsu"] = Vertex{40.68433, -74.39967}
	m["haiNan"] = Vertex{40.68433, -74.39967}
	m["tianJing"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	delete(m, "baijing")
	fmt.Println(m)
	m1, has := m["haiNan"]
	fmt.Println(m1, has)
	fmt.Println(m1.abs())

	f :=mystring("hello")
	fmt.Println(f.Abs())
}


