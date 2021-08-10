package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"springmars.com/grpc-server/cine"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9991"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("Failed to listen: %v \n", err)
	}
	server := grpc.NewServer()
	// 注册服务
	cine.RegisterService(server)

	fmt.Println("Listen on " + Address)
	server.Serve(listen)
}
