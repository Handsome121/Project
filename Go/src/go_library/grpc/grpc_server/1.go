package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "your-module-path/hello" // 导入生成的gRPC代码

)

type server struct{}

// 实现gRPC服务接口中的SayHello方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051") // 监听端口50051
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer() // 创建一个gRPC服务器

	// 将服务注册到gRPC服务器上
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Server started on port 50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
