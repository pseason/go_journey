package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建Customer实例方法，不带id
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户的信息,格式化的字符串
func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t %v\t %v\t %v\t %v\t %v\t", c.Id,
		c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
