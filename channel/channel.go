package main

import (
	"fmt"
	"time"
)

func sum(s []int, ch chan int)  { //求和
	sum := 0
	for _, v := range s{
		sum += v
	}
	ch <- sum //最后把结果推到通道
}


func main()  {
	//block()

   s := []int{7,2,5,2,1,5,-5,5-4}

   c := make(chan int)
   go sum(s[:len(s)/2],c) //前半部分求和 16  随机执行
   go sum(s[len(s)/2:],c) //后半部分求和 2

   x,y := <- c, <-c

   fmt.Println(x,y,x+y)
}


func  block()  {
	go func() {
		time.Sleep(1 * time.Hour)
	}()

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}