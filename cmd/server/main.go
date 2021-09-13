package main

import (
	"fibonachi_gRPC/pkg/api"
	"fibonachi_gRPC/pkg/fibonachi"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &fibonachi.GRPCServer{}
	api.RegisterFibonachiSliceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
