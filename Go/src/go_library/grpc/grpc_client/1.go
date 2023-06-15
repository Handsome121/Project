package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "your-module-path/hello" // 导入生成的gRPC代码
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // 连接到gRPC服务器
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn) // 创建gRPC客户端

	req := &pb.HelloRequest{Name: "John"} // 创建请求消息

	resp, err := client.SayHello(context.Background(), req) // 调用远程的gRPC服务
	if err != nil {
		log.Fatalf("Failed to say hello: %v", err)
	}

	log.Println("Response:", resp.GetMessage()) // 打印响应消息
}
