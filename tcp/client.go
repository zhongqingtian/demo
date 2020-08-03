package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var ch = make(chan int)

func TcpClient() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	coon, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("server is not starting")
		return
	}

	defer coon.Close()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("client send：%s", input)
		}
		//将从输入中读取的内容写入到连接中
		b := []byte(input)
		coon.Write(b)

		select {
		case <-ch:
			fmt.Println("server error，please reconnecting")
			return
		default:
			//不加default的话，那么<-ch会阻塞for，下一个输入就没法进行
		}
	}
}
