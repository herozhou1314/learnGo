package subchild

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const TimeFormat1 = "2006-01-02"

/***
*相同参数
 */
func add(x, y int) int {
	return x + y
}

/**返回两个参数**/
func swap(x, y string) (string, string) {
	return y, x
}

/**返回方法中的局部变量*/
func sprit(a int) (x, y int) {
	x = a * 10
	y = x * a
	return
}
func doSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
	fmt.Println("\nWhen's Saturday?")

	t := time.Now()
	fmt.Println(t.Year(), t.YearDay(), t.Day(), t.Hour())

	today := time.Now().Weekday()

	fmt.Println(today)
	fmt.Println(time.Now().Format(TimeFormat))
	t, err := time.Parse(TimeFormat, "2013-08-11 11:18:46")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(t.Format(TimeFormat1))

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

func doIf() {
	if os := runtime.GOOS; strings.EqualFold(os, "Linux") {
		fmt.Println("EqualFold")
	} else {
		fmt.Println(" not EqualFold")
	}
}
func main_() {
	fmt.Println(add(2, 2))
	a, b := swap("123", "456")
	fmt.Println(a, b)
	x, y := sprit(10)
	fmt.Println(x, y)
	c, d := 2, 3.45
	b1 := int(d)
	fmt.Println(c, d, b1)
	for i := 0; i < 20; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println("")
	j := 0
	for j < 15 {
		j++
	}
	fmt.Println(j)
	doSwitch()
	doIf()
}
