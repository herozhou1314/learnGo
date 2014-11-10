package main

import "fmt"
var syn chan int = make(chan int)
func main(){
	fmt.Println("hello")
	//printHello()
	go printHello()
	//printHello()
//	fmt.Println("I'm listening.")
//	time.Sleep(2 * time.Second)
//	fmt.Println("You're boring; I'm leaving.")
	<- syn
}

func printHello(){
	fmt.Println("printHello")
	syn <- 1
}
