package main

import (
	"fmt"
	"net/http"

	"springmars.com/http-mysql/repo"
	"springmars.com/http-mysql/web"
	_ "springmars.com/http-mysql/web"
)

func main() {
	// 初始化MysqlDB
	repo.InitDb()
	// 启动http服务
	listenAndServe()
	// panic 回调 recover
	defer recoverErr()
}

// listenAndServe
func listenAndServe() {
	// 循环注册
	for path, handler := range web.Handlers {
		http.HandleFunc(path, handler)
	}
	// 监听
	fmt.Println("Http ListenAndServe:", "192.168.50.188:9889")
	err := http.ListenAndServe("192.168.50.188:9889", nil)
	// err
	if err != nil {
		panic("程序启动失败:" + err.Error())
	}
}

// panic 异常退出可执行recover func
func recoverErr() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}
