package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro/grpc/hello"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Msg: "Hello World" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听失败：%v", err)
	}

	s := grpc.NewServer()
	hello.re(s, &server{})
	s.Serve(lis)
}
