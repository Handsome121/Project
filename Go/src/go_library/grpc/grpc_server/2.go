package main

import (
	pb "Go/go_library/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server1 struct {
}

type server struct {
	pb.UnimplementedGreetServer
	server1
}

func (s *server) Hello(ctx context.Context, req *pb.StreamReqData) (*pb.StreamResData, error) {
	return &pb.StreamResData{
		Data: req.Data + " from server",
	}, nil
}

func (s server1) Bye(ctx context.Context, req *pb.StreamReqData) (*pb.StreamResData, error) {
	return &pb.StreamResData{
		Data: req.Data + " from server",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &server{})

	fmt.Printf("Server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
