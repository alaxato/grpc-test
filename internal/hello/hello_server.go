package hello

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpc-test/proto"
	"time"
)

type HelloServer struct {
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (h *HelloServer) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*pb.GetHelloResponse, error) {
	fmt.Println("GetHello", req.PingDate.AsTime())
	return &pb.GetHelloResponse{
		Id:       req.Id,
		PongDate: timestamppb.New(time.Now()),
	}, nil
}
