package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Hight int
}

func main() {
	// 1.连接远程rpc服务
	client, err := rpc.DialHTTP("tcp", ":8998")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用方法
	ret := 0
	err = client.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Rect.Area:", ret)
	err = client.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Rect.Perimeter:", ret)
}
