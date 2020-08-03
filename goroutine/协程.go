package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
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