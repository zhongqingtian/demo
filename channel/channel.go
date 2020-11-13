package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(s []int, ch chan int) { //求和
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum //最后把结果推到通道
}

func main1() {
	//block()

	s := []int{7, 2, 5, 2, 1, 5, -5, 5 - 4}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //前半部分求和 16  随机执行
	go sum(s[len(s)/2:], c) //后半部分求和 2

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func block() {
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

func ChanLen() (int, int, int) {
	ch := make(chan int, 1)
	l1 := len(ch)
	ch <- 1
	l2 := len(ch)
	close(ch)
	l3 := len(ch)
	return l1, l2, l3
}

func bibao(f func() int) {
	f()
	f = nil
}

func BB() {
	//	var f *func() int
	t := time.NewTicker(1 * time.Second)
	t2 := time.NewTicker(time.Second * 2)
	var f1 func()
	defer func() {
		t.Stop()
		t2.Stop()
	}()
	for i := 0; i < 20; i++ {
		select {
		case <-t.C:
			d := &Demo{
				User: "user",
				Age:  0,
			}
			d.Age = rand.Intn(100)
			fmt.Println(d.Age)
			if d.Age < 50 {
				f1 = func() {
					k := d
					fmt.Println(k)
				}
			}
		case <-t2.C:
			if f1 != nil {
				f1()
			}
		}
	}
}

type Demo struct {
	User string
	Age  int
}
