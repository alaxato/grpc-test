package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpc-test/proto"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Second)

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewHelloServiceClient(conn)
	response, err := client.GetHello(ctx, &pb.GetHelloRequest{
		Id:       "1",
		PingDate: timestamppb.New(time.Now()),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID:", response.Id, "PongDate:", response.PongDate.AsTime())
}
