package main

import (
	"fmt"

	pb "demo/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	//连接指定的IPgRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	//第二步
	//初始化客户端
	c := pb.NewHelloClient(conn)

	//根据 pb里面json格式创建一个结构对象，然后存值
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC" //传入值

	//第三步
	// 调用服务器方法,本地是pb.go 的接口 ，接收返回的结构
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}

	//打印 返回数据
	fmt.Println(r.Message)

}
