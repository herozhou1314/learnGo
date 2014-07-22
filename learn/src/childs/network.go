package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("this is EmulateLoginBaidu.go\n")
	baiduMainUrl := "http://www.baidu.com/"
	fmt.Printf(baiduMainUrl)
	res, error := http.Get("http://bbs.golang-china.org/")
	if error != nil {
		panic(error)
	}
	fmt.Println(res)
}
