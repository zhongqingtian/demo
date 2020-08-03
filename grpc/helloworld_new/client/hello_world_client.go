package client

import (
	"context"
	pb "demo/grpc/helloworld_new/proto"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func StartClient() {

	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())

	c := pb.NewHelloServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HelloWorldClientAndServerStream(ctx, grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	for i := 0; i < 10; i++ { // 连续发送10个请求流
		r.Send(&pb.HelloRequest{Request: "my is golang gRpc client " + strconv.Itoa(i)})
	}
	r.CloseSend()
	for {
		res, err := r.Recv() //连续读取客户端发送过来的请求流数据，只有读取失败才停止循环
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			break
		}
		log.Printf("result:%v", res.Response)
	}
	defer conn.Close()
}
