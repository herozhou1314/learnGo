package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	p = Vertex{1, 2}
	q = &Vertex{1, 2}
	r = Vertex{X: 1} // Y:0 被省略
	s = Vertex{}     // X:0 和 Y:0
)

type User struct {
	name    string
	address string
}

func doArray() {
	fmt.Println("done the array")
	var arrays [2]string;
	arrays[0] = "11"
	arrays[1] = "12"
	fmt.Println(arrays)
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("p[%d] == %d\n", i, arr1[i])
	}
	fmt.Println(arr1[1:3])
	fmt.Println(arr1[:3])
	fmt.Println(arr1[4:])
}
func main() {
	v := Vertex{1, 2}
	v.X = 100
	v1 := &v
	v1.Y = 231
	fmt.Println(v)
	user := User{"herozhou1314", "street xili"}
	fmt.Println(user.name)
	fmt.Println(p, q, r, s)
	var Ui *User = new(User)
	Ui.name, Ui.address = "111", "111"
	fmt.Println(Ui)
	/*mas := map[string]User{{"1":User{"hero", "xili"}}}
	fmt.Println(mas)*/
	//	var arrays [2]User={{"hero","112300"},{"1234","rrrrrr"}}
	doArray()

}
