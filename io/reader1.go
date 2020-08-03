package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
// 读取器，它将数据从某个资源读取到传输缓冲区
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil{
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n,string(p[:n]))
	}
}