package main
import (
	"encoding/json"
	"fmt"
	"os"
)
type Server struct {
	ServerName string `json:serverName`
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}
type Server1 struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`
	// ServerName 的值会进行二次JSON编码
	ServerName  string //`json:"serverName"`
	ServerName2 string `json:"serverName2,string`
	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP   string `json:"serverIP,omitempty"`
}
func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	for _, arr := range s.Servers {
		fmt.Println(arr.ServerIP, "    "+arr.ServerName)
	}
	/**生成Json*/
	s1 := Server1 {
		ID:         3,
		ServerName:  `Go 1.0`,
		ServerName2: `Go 1.0`,
		ServerIP:   `123`,
	}
	b, _ := json.Marshal(s1)
	os.Stdout.Write(b)
	//	i, _ := js.Get("test").Get("int").Int()
	//ms := js.Get("test").Get("string").MustString()
}
