package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

// @Description 给算术对象 计算乘积的方法
// 参数 args是个 Args结构体，存储传入的计算参数
// 参数 reply 作为结果地址，存储最终结果
// return nil
func (t *Arith) Multiply(args *Args, reply *int) error {
	time.Sleep(time.Second * 3) //睡三秒，同步调用会等待，异步会先往下执行，用于客户端测试
	*reply = args.A * args.B
	return nil
}

// 计算商和余数
// 参数 args是个 Args结构体，存储传入的计算参数
// 参数 que 作为结果地址，存储最终结果
// return nil
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	time.Sleep(time.Second * 3)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//创建对象
	arith := new(Arith)
	//rpc服务注册了一个arith对象
	// 公开方法供客户端调用
	rpc.Register(arith) //必须要在rpc上面注册对象，暴露调用接口

	//配置指定rpc的传输协议 这里采用http协议作为rpc调用的载体 也可以用rpc.ServeConn处理单个连接请求
	rpc.HandleHTTP()
	//通过监听 tcp 127.0.0.1:1234 端口，建立连接后返回一个Listener
	k, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("listen error", e)
	}

	go http.Serve(k, nil)

	os.Stdin.Read(make([]byte, 1))
}
