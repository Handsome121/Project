/*
服务器流式
*/
package main

import (
	pb "Go/go_library/grpc/proto/hellopb/hellopb"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (*server) SayHello(req *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	for i := 0; i < 10; i++ {
		message := "Message " + strconv.Itoa(i)
		if err := stream.Send(&pb.HelloReply{Message: message}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
