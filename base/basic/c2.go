package main

import "fmt"

/**
变量声明
	标准声明： var 变量名 变量类型
	批量声明： var (
		a int
		b string
	)
	简短声明： 变量名 := 表达式
*/
func main() {
	// 标准声明
	var i int
	i = 10
	fmt.Println(i)
	// 批量声明
	var (
		a int
		b string
	)
	a = 5
	b = "many"
	fmt.Printf("%d - %s \n", a, b)
	// 简短声明
	c := "xxx"
	fmt.Println(c)
}
