package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

type MyMux struct{}
type test_struct struct {
	Test string
}
func (h MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r)
	 //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(r.Body)
	//fmt.Println(r.)
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	var h MyMux
	//http.HandleFunc("/", ServeHTTP()) //设置访问的路由
	err := http.ListenAndServe(":9092", h) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
