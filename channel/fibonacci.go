package channel

import "fmt"

func fibanacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { //不断地进行循环 费波纳数列
		case c <- x:
			x, y = y, x+y //通道有值进来 就加
		case <-quit: //直到 quit 接收
			fmt.Println("quit")
			return
		}
	}
}

func main2() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		} //数10个数后进行停止 向quit通道发送信号
		quit <- 0
	}()

	fibanacci(c, quit)
}
