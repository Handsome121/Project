package main

import (
	pb "Go/go_library/grpc/proto" //替换为你的包名
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"sync"
	"time"
)

type server struct {
	conn   *grpc.ClientConn
	client pb.GreeterClient
}

func (s *server) Start() {
	var err error

	s.conn, err = grpc.DialContext(context.Background(), "localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
	}

	s.client = pb.NewGreeterClient(s.conn)
}

// 单向模式
func (s *server) SayHello() {
	req := &pb.StreamReqData{Data: "hello"}

	resp, err := s.client.SayHello(context.Background(), req)
	if err != nil {
		if statusInfo, ok := status.FromError(err); ok {
			fmt.Println(statusInfo.Code())
			fmt.Println(statusInfo.Message())
		}
	}
	fmt.Println("Response:", resp)
}

// 客户端流模式
func (s *server) PutStream() {
	req := &pb.StreamReqData{Data: "hello"}

	putRes, _ := s.client.PutStream(context.Background())
	for i := 0; i < 10; i++ {
		req.Data = fmt.Sprintf("NO：%d, 客户端流模式", i)
		fmt.Println("向服务端发送", req.Data)
		err := putRes.Send(req)
		if err != nil {
			_ = fmt.Errorf("向服务端发送数据失败,err is %v", err)
		}
	}
}

//服务端流模式
func (s *server) GetStream() {
	req := &pb.StreamReqData{Data: "hello"}

	res, _ := s.client.GetStream(context.TODO(), req)
	for {
		data, err := res.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务端传输结束")
				break
			}
			_ = fmt.Errorf("err is %v", err)
		}
		fmt.Println(data)
	}
}

//双向流模式
func (s *server) AllStream() {
	req := &pb.StreamReqData{Data: "hello"}

	allStream, _ := s.client.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			StreamResData, _ := allStream.Recv()
			fmt.Println("AllStream收到服务端消息", StreamResData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		req.Data = "AllStream：我是客户端"
		for i := 0; i < 10; i++ {
			_ = allStream.Send(req)
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	var server server

	server.Start()
	server.SayHello()
	//server.PutStream()
	//server.GetStream()
	//server.AllStream()
}
