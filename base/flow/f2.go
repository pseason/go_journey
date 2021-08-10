/*
2. Go语言for（循环结构）
	a. 使用循环语句时，需要注意的有以下几点：
		- 左花括号{必须与 for 处于同一行。
		- Go语言中的 for 循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，
			Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量。
		- Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break
*/

package main

import "fmt"

// for 基础用法
func forBasic() {
	sum := 0
	for {
		sum++
		if sum > 99 {
			break
		}
	}
	fmt.Printf("sum value: %d \n", sum)

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d \n", i)
	}
	// 乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d	", i, j, i*j)
		}
		fmt.Println("")
	}
	fmt.Println("-------------")
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d	", i, j, i*j)
		}
		fmt.Println("")
	}
}

// for break 特殊用法
func forSpecial() {
	sum := 0
JLoop:
	for {
		if sum > 99 {
			break JLoop
		}
		sum++
	}
	fmt.Printf("sum value: %d", sum)
}
func main() {
	forBasic()
	forSpecial()
}
