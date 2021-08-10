package main

import (
	"springmars.com/customer/service"
	"springmars.com/customer/view"
)

func main() {
	//在main函数中，创建一个customerView,并运行显示主菜单..
	customerView := view.CustomerView{
		Key:  "",
		Loop: true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.CustomerService = service.NewCustomerService()
	//显示主菜单..
	customerView.MainMenu()
}
