package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "Go/go_library/grpc/proto" // 替换为你的包名

)

type server struct {
	pb.UnimplementedGreeterServer
}

// 单相模式
func (s *server) SayHello(ctx context.Context, req *pb.StreamReqData) (*pb.StreamResData, error) {
	fmt.Println("server SayHello run")
	//return &pb.StreamResData{
	//	Data: req.Data + " from server",
	//}, nil
	return nil,status.Error(codes.Internal,"internal error")
}

// 客户端流模式
func (s *server) PutStream(clientData pb.Greeter_PutStreamServer) error {
	fmt.Println("server PutStream run")
	for {
		a, err := clientData.Recv()
		if err != nil {
			_ = fmt.Errorf("err is %v", err)
			break
		}
		fmt.Println(a)
	}

	return nil
}

// 服务端流模式
func (s *server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {
	fmt.Println("server GetStream run")
	for i := 0; i < 10; i++ {
		err := res.Send(&pb.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			_ = fmt.Errorf("send err,err is %v", err)
		}
		time.Sleep(time.Second)
	}
	return nil
}

// 双向流模式
func (s *server) AllStream(allData pb.Greeter_AllStreamServer) error {
	fmt.Println("server AllStream run")
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data, _ := allData.Recv()
			fmt.Println("AllStream收到客户端消息", data)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			_ = allData.Send(&pb.StreamResData{
				Data: "我是服务器",
			})
			time.Sleep(time.Second)
		}
	}()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	fmt.Printf("Server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
