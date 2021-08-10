package main

import (
	"fmt"
	"math"
)

/**
字符格式化输出
	%v	按值的本来值输出
	%+v	在 %v 基础上，对结构体字段名和值进行展开
	%#v	输出 Go 语言语法格式的值
	%T	输出 Go 语言语法格式的类型和值
	%%	输出 % 本体
	%b	整型以二进制方式显示
	%o	整型以八进制方式显示
	%d	整型以十进制方式显示
	%x	整型以十六进制方式显示
	%X	整型以十六进制、字母大写方式显示
	%U	Unicode 字符
	%f	浮点数
	%p	指针，十六进制方式显示
*/

func main() {
	var s string = "在 %v 基础上，对结构体字段名和值进行展开"
	f := fmt.Sprintf(s, "y")
	fmt.Println(f)
	ff := fmt.Sprintf("%v %v %v", math.Pi, "数字", 1)
	fmt.Println(ff)

}
