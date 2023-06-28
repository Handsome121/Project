/*
服务器流式
*/
package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "Go/go_library/grpc/proto/hellopb/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 调用服务器流式传输方法
	req := &pb.HelloRequest{}
	stream, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call StreamData: %v", err)
	}

	// 接收服务器发送的消息
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive stream response: %v", err)
		}
		log.Printf("Received message: %s", resp.Message)
	}
}
