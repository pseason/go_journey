package main

import (
	"fmt"
	"net/http"

	"springmars.com/http-mysql-viper/config"
	"springmars.com/http-mysql-viper/repo"
	"springmars.com/http-mysql-viper/web"
	_ "springmars.com/http-mysql-viper/web"
)

func main() {
	// 加载配置
	config.LoadConfig()
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
	fmt.Println("Http ListenAndServe:", "192.168.50.188", config.AppConf.Port)
	err := http.ListenAndServe(fmt.Sprintf("192.168.50.188:%d", config.AppConf.Port), nil)
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
