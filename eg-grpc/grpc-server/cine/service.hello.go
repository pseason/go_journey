package cine

import (
	"context"
	"fmt"

	"springmars.com/grpc-cine/icine"
)

type HelloService struct {
}

func (h HelloService) SayHello(ctx context.Context, in *icine.HelloReq) (*icine.HelloResp, error) {

	fmt.Println("be invoked", in.Name)

	resp := new(icine.HelloResp)
	resp.Msg = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}
