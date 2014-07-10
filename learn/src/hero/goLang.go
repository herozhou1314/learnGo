package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func printFile() {
	f, err := os.Create("haha2.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"编号", "姓名", "年龄"})
	w.Write([]string{"1", "张三", "23"})
	w.Write([]string{"2", "李四", "24"})
	w.Write([]string{"3", "王五", "25"})
	w.Write([]string{"4", "赵六", "26"})
	w.Flush()
}

func sum(arg []int) int {
	a := 0
	for i := 0; i < len(arg); i++ {
		a += arg[i]
	}
	return a
}



func main() {
	fmt.Println("Hello, 世界")

	str := "hello  world"

	fmt.Println(str)

	var a, b, c = 1, "hello", true

	fmt.Println(a, strings.Index(b, "o"), c)

	for i := 0; i < 10; i++ {
		fmt.Println("count %d-->>", i)
	}
	const key = "my name is herozhou"
	fmt.Println(key)

	s := "hello"
	s = "c"+s[3:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
	s12 := `hello word`
	fmt.Println(s12)
	printFile()
	m := map[string]int{"one": 1, "two": 2}

	m1 := map[int]int{1: 1, 2: 2}

	fmt.Println(m)
	fmt.Println(m1)

	s2 := sum([]int{1, 2, 3})

	fmt.Println(s2)

	//maps := map[string]string{{"hero": "baby"}, {"key": "value"}}
	//	fmt.Println(maps)
	fmt.Println("hello go vim")
	//childs.doChild();
	map2 := map[string]string{{"1":"111"}, {"2":"2222"}}
	fmt.Println(map2)
}
