package main

import (
	pb "activity/contract"
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetBody())
	return &pb.Response{Message: in.GetBody()}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
