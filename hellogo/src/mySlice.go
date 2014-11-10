package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello", time.Now())
	fmt.Println("---------------------------------------------------")

	a := make([]int, 5)
	arr1 := make([]string, 5)
	arr1[0] = "123"
	arr1[1] = "123"
	arr1 = append(arr1, "hello")
	arr1 = append(arr1, "hello")
	printSlice1("arr1", arr1)

	fmt.Println(arr1)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice1(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
