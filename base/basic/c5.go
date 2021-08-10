package main

import "fmt"

/**
变量作用域
	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数

*/

func part() {
	// 局部变量
	a, b := 10, "s"
	fmt.Println(a, b)
}

func append(to string) string {
	return s + " " + to
}

var s string = "hello"

func main() {
	//局部
	part()
	//全局
	s = s + " " + "world"
	fmt.Println(s)
	// 形参
	fmt.Println(append(", go"))
}
