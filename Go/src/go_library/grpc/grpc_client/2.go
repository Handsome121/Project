package main

import (
	pb "Go/go_library/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct {
	conn   *grpc.ClientConn
	client pb.GreetClient
}

func (s *server) Start() {
	var err error

	s.conn, err = grpc.DialContext(context.Background(), "localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
	}

	s.client = pb.NewGreetClient(s.conn)
}

// 单向模式
func (s *server) SayHello() {
	req := &pb.StreamReqData{Data: "hello"}

	resp, err := s.client.Hello(context.Background(), req)
	if err != nil {
		if statusInfo, ok := status.FromError(err); ok {
			fmt.Println(statusInfo.Code())
			fmt.Println(statusInfo.Message())
		}
	}
	fmt.Println("Response:", resp)
}
