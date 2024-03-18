package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/internal/hello"
	pb "grpc-test/proto"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, hello.NewHelloServer())
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		return
	}
}
