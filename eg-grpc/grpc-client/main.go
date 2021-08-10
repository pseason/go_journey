package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"springmars.com/grpc-cine/icine"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9991"
)

func main() {

	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := icine.NewGreeterClient(conn)
	// 参数
	req := &icine.HelloReq{Name: "gRPC"}

	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Msg)
}
