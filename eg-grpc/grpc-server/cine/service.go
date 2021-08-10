package cine

import (
	"google.golang.org/grpc"

	"springmars.com/grpc-cine/icine"
)

// register 服务
func RegisterService(server *grpc.Server) {
	icine.RegisterGreeterServer(server, HelloService{})
}
