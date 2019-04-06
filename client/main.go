package main

import (
	"context"
	"fmt"

	"github.com/yoshd/go-grpc-lt/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:13009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewSampleClient(conn)
	req := pb.HelloRequest{Name: "Yoshd"}
	res, err := c.Hello(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
