/*
 Go语言type关键字（类型别名）
	1. 类型定义：
            使用方法 type name 基础类型 -> type Type int
	2. 类型别名：
		使用方法 type name = 基础类型 -> type Alias = int
*/
package main

import (
	"fmt"
)

// 类型定义
type Pay int

const (
	Bank Pay = iota + 1
	WeiXin
	AliPay
)

// 别名定义
type Car = int

const (
	BYD Car = iota + 1
	Benz
	BMW
	Audi
)

type cc struct {
}

func xx(inter interface{}) {
	fmt.Printf("xxx: %p , xxx: %v \n", &inter, inter)
}

func main() {
	fmt.Println(Bank, WeiXin, AliPay)
	fmt.Println(BYD, Benz, BMW, Audi)

	var MeiTuan Pay
	var XXL Car
	fmt.Printf("%T \n", MeiTuan)
	fmt.Printf("%T \n", XXL)

	var c cc
	fmt.Printf("xxx: %p , xxx: %v \n", &c, &c)
	xx(&c)
}
