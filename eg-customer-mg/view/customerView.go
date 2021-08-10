package view

import (
	"fmt"

	"springmars.com/customer/model"
	"springmars.com/customer/service"
)

type CustomerView struct {
	//定义必要字段
	Key  string //接收用户输入...
	Loop bool   //表示是否循环的显示主菜单
	//增加一个字段CustomerService
	CustomerService *service.CustomerService
}

//显示所有的客户信息
func (c *CustomerView) list() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := c.CustomerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (c *CustomerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if c.CustomerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

//得到用户的输入id，删除该id对应的客户
func (c *CustomerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Print("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if c.CustomerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出软件
func (c *CustomerView) exit() {
	fmt.Print("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&c.Key)
		if c.Key == "Y" || c.Key == "y" || c.Key == "N" || c.Key == "n" {
			break
		}
		fmt.Print("你的输入有误，确认是否退出(Y/N)：")
	}
	if c.Key == "Y" || c.Key == "y" {
		c.Loop = false
	}
}

//显示主菜单
func (c *CustomerView) MainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&c.Key)
		switch c.Key {
		case "1":
			c.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !c.Loop {
			break
		}
	}
	fmt.Println("已退出了客户关系管理系统...")
}
