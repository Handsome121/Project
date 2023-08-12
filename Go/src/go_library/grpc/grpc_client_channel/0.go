package grpc_client_channel

import pb "Go/go_library/grpc/proto"

var GrpcClientChan chan pb.GreeterClient

const (
	GrpcClientChanSize int = 1
)

func init() {
	GrpcClientChan = make(chan pb.GreeterClient, GrpcClientChanSize)
}
