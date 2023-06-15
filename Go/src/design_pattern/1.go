/*
代理模式
*/
/*
如果你将这些创建的GRPC服务端看作是一个池子，并且只有一个池子出口用于与外部系统对接，你可以采用代理模式来实现。
代理模式可以将多个服务端封装在一个代理对象中，对外部系统提供统一的接口。该代理对象负责管理服务端池子，接收外部系统的请求，并将请求分发给空闲的服务端进行处理。
*/
package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// 自定义GRPC服务
type MyService struct{}

// 定义GRPC方法
func (s *MyService) MyMethod(ctx context.Context, request *MyRequest) (*MyResponse, error) {
	// 在这里处理你的业务逻辑
	// request是输入参数，MyResponse表示输出数据，error表示错误（如果有）

	// 示例：简单返回一个固定的响应
	response := &MyResponse{
		Message: "Hello, client!",
	}
	return response, nil
}

// 服务端池子代理对象
type ServerPoolProxy struct {
	servers []*grpc.Server
	mu      sync.Mutex
	index   int
}

// 初始化服务端池子
func NewServerPoolProxy(numServers int) *ServerPoolProxy {
	pool := &ServerPoolProxy{
		servers: make([]*grpc.Server, numServers),
		index:   0,
	}

	for i := 0; i < numServers; i++ {
		server := grpc.NewServer()
		myService := &MyService{}
		RegisterMyServiceServer(server, myService)
		pool.servers[i] = server
	}

	return pool
}

// 处理客户端请求
func (p *ServerPoolProxy) HandleRequest(request *MyRequest) (*MyResponse, error) {
	p.mu.Lock()
	server := p.servers[p.index]
	p.index = (p.index + 1) % len(p.servers)
	p.mu.Unlock()

	conn, err := net.Dial("tcp", ":50051") // 假设服务端监听的端口为50051
	if err != nil {
		return nil, err
	}

	client := NewMyServiceClient(server)
	response, err := client.MyMethod(context.Background(), request)
	conn.Close()

	return response, err
}

func main() {
	// 创建服务端池子代理对象
	serverPool := NewServerPoolProxy(3) // 假设池子中有3个服务端

	// 模拟外部系统请求
	request := &MyRequest{
		Data: "Hello, server!",
	}

	// 处理外部系统请求
	response, err := serverPool.HandleRequest(request)
	if err != nil {
		log.Fatalf("Failed to handle request: %v", err)
	}

	log.Printf("Response: %s", response.Message)
}







