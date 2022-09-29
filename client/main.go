package main

import (
	"context"
	"fmt"
	"grpc-demo/hellogrpc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := fmt.Sprintf("localhost:%d", 8989)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hellogrpc.NewHelloServiceClient(conn)
	resp, err := client.SayHello(context.Background(), &hellogrpc.HelloRequest{Greeting: "大刘呀"})
	if err != nil {
		log.Fatal("调用RPC 方法错误:", err)
	}
	fmt.Println("调用rpc 方法成功,返回结果：", resp.Reply)

}
