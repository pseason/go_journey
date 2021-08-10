package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"springmars.com/gin-mysql-viper/config"
	"springmars.com/gin-mysql-viper/repo"
	"springmars.com/gin-mysql-viper/web"
	_ "springmars.com/gin-mysql-viper/web"
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
	engine := gin.Default()
	// 跨域
	engine.Use(web.Cors())
	// 循环注册
	for _, route := range web.Handlers {
		switch route.Method {
		case web.GET:
			engine.GET(route.Url, route.Handler)
		case web.POST:
			engine.POST(route.Url, route.Handler)
		case web.DELETE:
			engine.DELETE(route.Url, route.Handler)
		case web.PUT:
			engine.PUT(route.Url, route.Handler)
		case web.PATCH:
			engine.PATCH(route.Url, route.Handler)
		default:
			fmt.Println("路由注册失败:", route.Method, route.Url)
		}
	}

	fmt.Println("Http ListenAndServe:", "192.168.50.188", config.AppConf.Port)
	if err := engine.Run(fmt.Sprintf("192.168.50.188:%d", config.AppConf.Port)); err != nil {
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
