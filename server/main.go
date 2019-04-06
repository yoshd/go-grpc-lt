package main

import (
	"context"
	"fmt"
	"net"

	"github.com/yoshd/go-grpc-lt/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello %s!", in.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:13009")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterSampleServer(s, &server{})
	s.Serve(lis)
}
