package main
import (
	"fmt"
	"math"
	"runtime"
	"time"
)
var m, n int = 1, 5
var s, d string = "a", "dd"
func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello World!", math.Pi)
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World!", i)
	}
	a, b := add(1, 5, "a", "v")
	fmt.Println(a, b)
	x, y := add2(1, 5, "hello ", "world");
	fmt.Println(x, y)
	fmt.Println(m, n)
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	for i := 0; i < 100; i++ {
	}
	i := 0
	for i < 100 {
		i++;
	}
	if a := 10; a < 5 {
		fmt.Println("hello")
	}else if a > 100 {
		fmt.Println("hello---->>")
	}else {
		fmt.Println("hello--wwww")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	time1();
}
func add(x , y int, a , b string) (int, string ) {
	return (x+y), a+b
}
func add2(x, y int , str1, str2 string) (a int , b string) {
	a = x+y
	b = str1+str2
	return
}
func time1(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
