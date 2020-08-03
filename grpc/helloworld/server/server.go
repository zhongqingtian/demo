package main

import (
	"context"
	pb "demo/grpc/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//在grpc中的定义 type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
	// 声明一个变量 其中一元拦截器只能设置拦截一个服务
	var serverInterceptor grpc.UnaryServerInterceptor
	// 为这个变量赋值
	serverInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 这个方法的具体实现
		err = check(ctx) //权限校验
		if err != nil {
			return
		}
		// 校验通过后继续处理请求
		return handler(ctx, req)
	}
	// 将拦截器添加进去
	s := grpc.NewServer([]grpc.ServerOption{grpc.UnaryInterceptor(serverInterceptor)}...)

	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("server start successful")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func check(ctx context.Context) error {
	return nil
}

// 这个方法接受不多个拦截器，最终返回一个拦截器
func InterceptChain(intercepts ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	// 获取拦截器的长度
	l := len(intercepts)
	// 如下我们返回一个拦截器
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 在这个拦截器中，我们做一些操作
		//构造一个链
		chain := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(
					currentCtx,
					currentReq,
					info,
					currentHandler)
			}
		}
		// 声明一个hander
		chainHandler := handler
		for i := l - 1; i >= 0; i-- {
			// 递归一层层调用
			chainHandler = chain(intercepts[i], chainHandler)
		}
		return chainHandler(ctx, req)
	}
}
