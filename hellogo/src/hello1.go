package main
import (
	"fmt"
	"time"
)

func main() {
	var arr [2]string
	arr[0]="hello"
	arr[1]="world"
	p:=&arr
	fmt.Println("Hello World!",p)
	fmt.Println("Hello World!")
	fmt.Println("Hello World!",time.Now().Hour())

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


