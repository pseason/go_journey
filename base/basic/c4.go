package main

import "fmt"

/**
匿名变量

*/

func calulate(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	// 函数返回值全量接收
	a, b := calulate(100, 50)
	fmt.Println(a, b)
	// 函数返回值选择接收
	c, _ := calulate(1100, 50)
	fmt.Println(c)
}
