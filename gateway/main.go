package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/yoshd/go-grpc-lt/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterSampleHandlerFromEndpoint(ctx, mux, "localhost:13009", opts)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", mux)
}
