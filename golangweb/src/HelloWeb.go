package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
	"io"
)
type jsonBody struct {
	Username   string
	Password   string
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println("==========")  //这些信息是输出到服务器端的打印信息
	fmt.Println(r.Header)  //这些信息是输出到服务器端的打印信息
	fmt.Println(r.PostForm)  //这些信息是输出到服务器端的打印信息
	fmt.Println(r.Body)  //这些信息是输出到服务器端的打印信息
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, string(result)) //这个写入到w的是输出到客户端的
	r.Body.Close()
	fmt.Printf("%s\n", result)
	var s jsonBody
	json.Unmarshal(result, &s)
	fmt.Println(s)
	var f interface{}
	err := json.Unmarshal(result, &f)
	//断言的方式
	m := f.(map[string]interface{})
	fmt.Println("userName--->>",m["userName"])
	if err == nil {
		for k, v := range m {

			fmt.Println(k,v)

		}
	}
	io.WriteString(w, "\n hello, golang!\n")
}
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
