package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func Input(key string, mvp map[string]string) {
	mvp[key] = key
}
func OutPut(mvp map[string]string) {
	for s, s2 := range mvp {
		fmt.Println(s, s2)
	}
}
func main() {
	//go say("world") //开一个新的Goroutines执行
	//say("hello")    //当前Goroutines执行
	mvp := make(map[string]string)
	for i := 0; i < 10; i++ {
		go Input(fmt.Sprintf("%v", i), mvp)
	}
	for i := 0; i < 10; i++ {
		go Input(fmt.Sprintf("%v", i), mvp)
	}
	go OutPut(mvp)

	time.Sleep(3 * time.Second)
}

// 以上程序执行后将输出：
// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello
