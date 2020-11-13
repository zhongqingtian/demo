package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	//默认是无缓冲 阻塞的
	c := make(chan int)
	defer close(c)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	ch := make(chan int, 2) //加了数值，表示有缓冲，非阻塞的
	ch <- 1
	ch <- 10
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 利用range 取数据
	ch2 := make(chan int, 10)
	go fibonacci(cap(ch2), ch2)
	for i := range ch2 { //迭代
		fmt.Println(i)
	}
}
