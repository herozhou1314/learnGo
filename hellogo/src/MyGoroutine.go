package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50)
	go printHello(c)
	a := []int{7, 2, 8, -9, 4, 0}
	//go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	go func(a string, b int) {
		time.Sleep(3 * time.Second)
		fmt.Println("printHello", a, b)
		c <- 5 // 将和送入 c
	}("hello", 123)

	x, ok := <-c
	y, ok1 := <-c  // 从 c 中获取
	y1, ok2 := <-c // 从 c 中获取
	fmt.Println(x, y, x+y, ok, ok1, y1, ok2)

}

func printHello(c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("printHello")
	c <- (-100)
}
func sum(a []int, c chan int) {
	time.Sleep(1 * time.Second)
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("sum")
	c <- sum // 将和送入 c
	//close(c)
}
