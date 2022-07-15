package main

import (
	"fmt"
	"github.com/lycclsltt/system"
	"net/http"
	"time"
)

/*
覆盖如下几类:

cpu
内存
io状态
磁盘占用
负载
内、外网卡的流入、流出流量等
*/
func main() {
	go func() {
		for i := 0; i < 10000; i++ {
			go Get()
		}
	}()
	mem := &system.Mem{}
	mem.Collect()
	println("used:", mem.MemUsedFunc(""), "kb")
	println("free:", mem.MemFreeFunc(""), "kb")

	for true {
		net := &system.NetWork{}
		net.Collect()
		println(net.EthInMaxUseRateFunc(""), "kb")
		println(net.EthOutMaxUseRateFunc(""), "kb")
		time.Sleep(200*time.Nanosecond)
	}
}

func Get()  {
	resp,err := http.Get("https://www.baidu.com")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(resp.ContentLength)
}