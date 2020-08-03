package main

import (
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second): //这个函数是一个chan 等待 5s后 这个函数会自动读出一个数
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
}
