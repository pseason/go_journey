package main

import (
	"errors"
	"fmt"
	"time"
)

func rpcServer(ch chan string) {

	for {
		// 接收
		data := <-ch
		fmt.Println("server received:", data)

		time.Sleep(time.Second * 2)

		// 返回
		ch <- "roger"
	}
}

func rpcClient(ch chan string, msg string) (string, error) {
	// 发送消息
	ch <- msg
	select {
	// 正常接收
	case ack := <-ch:
		return ack, nil
	// 超时
	case <-time.After(time.Second):
		return "", errors.New("接收超时")
	}

}

func main() {
	ch := make(chan string)
	// 开启服务
	go rpcServer(ch)

	recv, err := rpcClient(ch, "hello")
	if err == nil {
		fmt.Println("返回：", recv)
	} else {
		fmt.Println("错误：", err)
	}
}
