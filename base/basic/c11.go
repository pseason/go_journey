package main

import (
	"fmt"
)

/**
go 语言常量和const关键字
	1. 定义常量的表达式必须为能被编译器求值的表达式，并且常量只能是布尔型、
			数字型（整数型、浮点型和复数）和字符串型。
	2. 常量的声明语法为：
		const name [type] = value\express
		const name = value\express
		const (
			name [type] = value
			name  = value
		)
		[type] 可不书写，可由编译器根据值推断类型
	3. iota 常量生成器
		常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，
		但是不用每行都写一遍初始化表达式。在一个 const 声明语句中，
		在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加1
*/

const a, b int = 8, 9
const c = 8 + 9
const str string = "5446"
const (
	d          = 45
	msg string = "xxxxxxxxx"
)

// iota 常量生成器
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type myString string

func main() {
	// 常量不可修改
	fmt.Println(a, b, c, str, d, msg)
	fmt.Println(Sunday, Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday)

	var s2 Weekday = 9
	fmt.Println(int(s2))
}
