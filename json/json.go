package json

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName,omitempty"` //omitempty  如果该字段为空，json格式不会输出
	ServerIP   string `json:"serverIp"`
}

type Serverslice struct {
	Servers []Server
}

func know() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s) //解析后的值 存放在s结构体中（前提是知道上面两个结构体）
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)
}

func unknow() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{} //定义一个接口
	err := json.Unmarshal(b, &f)

	/* 现在f接口被赋值 是下面一整个结构，但是我们不知道怎样
	  通过类型断言来取出
	f = map[str]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}*/

	if err != nil {
		panic("json err")
	}

	//通过类型断言 返回整个 map的类型 f是一个接口，不能直接被迭代，只是地址指向了 这个结构体
	m := f.(map[string]interface{}) //有这个类型就返回类型值
	fmt.Println(m)
	for k, v := range m { //迭代出 m 的value还是interface 所以还要类型断言一次
		switch vv := v.(type) { //type 获取类型
		case string:
			fmt.Println(k, "is str", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}

func toJson() {
	var s Serverslice
	var k []Server
	k = append(k, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	//	k = append(k,Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s) //返回的是 []byte 要转换成string
	c, _ := json.Marshal(k)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	fmt.Println(string(c))
}
