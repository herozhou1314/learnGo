package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func (v *Vertex) ads() float64 {
	return v.Long * v.Lat
}

/**go 里面的foreach   用关键字range 而且有 index*/
func doForRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for n, m := range pow {
		fmt.Printf("2**%d = %d\n", n, m)
	}
	//丢弃index
	for _, m := range pow {
		fmt.Printf("%d\n", m)
	}
}

func doneMaps() {
	m := map[string]Vertex{"baidu": Vertex{40.68433, -74.39967}, "google": Vertex{37.42202, -122.08408}}
	m["gaode"] = Vertex{40.68433111, -74.3996711}

	fmt.Println(m)
	m["gaode"] = Vertex{40.684333, -74.39}
	fmt.Println(m)
	delete(m, "gaode")
	//返回值和检查是否存在
	v, has := m["hello"]

	fmt.Println("v-->", v, has)
	fmt.Println(m["baidu"])
	fmt.Println(m)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	fmt.Println("hello world")
	doForRange()
	doneMaps()
	pos, neg := adder(), adder()
	fmt.Println(pos(1), pos(1))
	fmt.Println(pos(1), pos(1))
	fmt.Println(pos(1), pos(1))
	fmt.Println(neg(1), neg(1))
	add := func(x, y int) int {
		return x * y
	}

	fmt.Println(add(3, 4))

	ve := new(Vertex)
	ve.Lat = 100
	ve.Long = 1
	fmt.Println(ve.ads())
}
