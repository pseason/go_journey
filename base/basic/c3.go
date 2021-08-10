package main

import "fmt"

/**
变量初始化
	标准格式：      var 变量名 变量类型 = 表达式
	编译器推导：    var 变量名 = 表达式
	声明并初始化:   变量名 := 表达式
*/
func main() {
	// 标准格式
	var a int = 3 + 5
	fmt.Println(a)
	// 编译器推导
	var b = 8 + 9
	fmt.Println(b)
	// 声明并初始化
	c := 5 + 8
	fmt.Println(c)

	// 变量赋值传递
	var d int = 100
	var e int = 200
	d, e = e, d
	fmt.Println(d, e)
}
