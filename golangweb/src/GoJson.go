package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	for _, arr := range s.Servers {
		fmt.Println(arr.ServerIP, "    "+arr.ServerName)
	}

	js, err := simplejson.NewJson([]byte(str))
	if err!=nil {

	}
	arr, _ := js.Get("servers").Array()

	fmt.Println(arr)
//	i, _ := js.Get("test").Get("int").Int()
	//ms := js.Get("test").Get("string").MustString()
}
