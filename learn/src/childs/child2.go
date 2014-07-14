package main

import (
	"fmt"
)

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
	var arrays [2]string
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

/**go 语言中类似于java list 的东东**/
func doslice() {
	a := [5]string{"hello"}
	fmt.Println(len(a), cap(a))
	v := []string{}
	v1 := make([]string, 5, 9)
	v2 := make([]int, 5)

	fmt.Println("v2--value", v2)
	fmt.Println(len(v1), cap(v1))
	fmt.Println(len(v2), cap(v2))
	v = append(v, "hello", "world")
	fmt.Println(len(v), cap(v))
	fmt.Println(v)

}
func doMaps() {
	m := map[string]string{"1": "123", "2": "234"}
	fmt.Println(m)
	mas := map[string]User{"1": {"hero", "xili"}}
	fmt.Println(mas)
	arrays := []User{User{name: "hero", address: "112300"}, User{"1234", "rrrrrr"}}
	u := new(User)
	u.name = "123"
	u.address = "123"
	fmt.Println(u)
	fmt.Println(arrays)
	users := make([]User, 5)
	for i := 0; i < len(users); i++ {
		u12 := users[i]
		u12.address = fmt.Sprintf("adr-->>%d", i)
		u12.name = fmt.Sprintf("name-->>%d", i)
		users[i] = u12
	}
	fmt.Println(users)
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
	doArray()
	doslice()
	doMaps()
}
