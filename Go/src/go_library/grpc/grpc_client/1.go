package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "Go/go_library/grpc/proto" // 替换为你的包名
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{Name: "Alice"}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response: %s", resp.Message)
}
