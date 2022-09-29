package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-demo/hellogrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RpcHelloServiceServer struct {
	*hellogrpc.UnimplementedHelloServiceServer
}

func (svc RpcHelloServiceServer) SayHello(ctx context.Context, r *hellogrpc.HelloRequest) (*hellogrpc.HelloResponse, error) {

	res := &hellogrpc.HelloResponse{}
	res.Reply = "我收到的请求是：" + r.Greeting
	return res, nil
}

func (svc RpcHelloServiceServer) LotsOfReplies(*hellogrpc.HelloRequest, hellogrpc.HelloService_LotsOfRepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfReplies not implemented")
}

func (svc RpcHelloServiceServer) LotsOfGreetings(hellogrpc.HelloService_LotsOfGreetingsServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfGreetings not implemented")
}

func (svc RpcHelloServiceServer) BidiHello(hellogrpc.HelloService_BidiHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiHello not implemented")
}

// missing method mustEmbedUnimplementedHelloServiceServer
func main() {
	host := fmt.Sprintf("localhost:%d", 8989)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("lister err :%v", err)
	}
	grpcServer := grpc.NewServer()

	//rpcHello := RpcHelloServiceServer{}

	hellogrpc.RegisterHelloServiceServer(grpcServer, &RpcHelloServiceServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	fmt.Println("启动服务成功：" + host)
}
