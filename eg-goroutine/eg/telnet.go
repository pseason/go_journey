package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func server(address string, exit chan int) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("telnet 开启失败", address)
		exit <- -1
	}
	fmt.Println("telnet 开启成功", address)

	defer l.Close()

	// 循环监听
	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("telnet 开启会话失败")
			continue
		}

		// 根据连接开启会话, 这个过程需要并行执行
		go handleSession(conn, exit)
	}
}

// 处理会话
func handleSession(conn net.Conn, exit chan int) {
	fmt.Println("session 开启成功")

	reader := bufio.NewReader(conn)

	// 循环接收数据
	for {
		recv, err := reader.ReadString('\n')
		if err == nil {
			recv = strings.TrimSpace(recv)

			if !processRecv(recv, exit) {
				fmt.Println("session 接收消息 -> session close", recv, err)
				conn.Close()
				break
			}
			conn.Write([]byte("咨询：" + recv + "\r\n"))
		} else {
			fmt.Println("session 接收消息错误 session close", err)
			conn.Close()
			break
		}

	}

}

// 处理消息
func processRecv(recv string, exit chan int) bool {
	if strings.HasPrefix(recv, "@close") {
		fmt.Println("recv: Session closed")
		// 告诉外部需要断开连接
		return false
		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(recv, "@shutdown") {
		fmt.Println("recv: Server shutdown")

		exit <- 0
		// 告诉外部需要断开连接
		return false
	}
	// 打印输入的字符串
	fmt.Println("recv: ", recv)
	return true
}

func main() {
	// 创建一个程序结束码的通道
	exit := make(chan int)
	go server("127.0.0.1:8992", exit)

	code := <-exit

	os.Exit(code)
}
