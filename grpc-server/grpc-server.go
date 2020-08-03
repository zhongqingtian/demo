package main

import (
	pb "demo/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	//gRPC服务地址
	Address = "127.0.0.1:50052"
)

//定义一个helloServer并实现约定的接口
type helloService struct{}

//给helloServer添加一个方法
// 当客户端调用这个方法时 会执行 这个实现的方法体
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)             //创建一个结构体
	resp.Message = "hello" + in.Name + "." //存储值
	return resp, nil                       //返回结构体
}

var HelloServer = helloService{}

func main() {
	//第一步
	//监听 IP地址 是否进行连接
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}

	//第二步
	//实现gRPC Server ，返回一个server句柄
	s := grpc.NewServer()

	//第三步
	//注册helloServer为客户端提供服务，暴漏方法接口
	pb.RegisterHelloServer(s, HelloServer) //内部调用了s.RegisterServer()
	fmt.Println("Listen on" + Address)

	s.Serve(listen)

}

// 链式拦截器
func ChainUnaryServer(inInterceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	interceptors := make([]grpc.UnaryServerInterceptor, 0)
	// filter nil interceptor
	for _, itt := range inInterceptors {
		if itt != nil {
			interceptors = append(interceptors, itt)
		}
	}

	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler grpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				if curI == lastI {
					return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := interceptors[curI](currentCtx, currentReq, info, chainHandler)
				curI--
				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
