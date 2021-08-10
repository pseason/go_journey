package main

import "fmt"

/**
Go语言模拟枚举（const和iota模拟枚举）
	1. Go语言现阶段没有枚举类型
	2. 可以使用 const 常量配合 iota 来模拟枚举类型
*/

type EnumInt int

const (
	a EnumInt = iota
	b
	c
)

const (
	// iota 使用了一个移位操作，每次将上一次的值左移一位（二进制位），以得出每一位的常量值
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

// 声明芯片类型
type ChipType int

// 声明芯片类型
const (
	None ChipType = iota
	CPU
	GPU
)

//定义 ChipType 类型的方法 String()，返回值为字符串类型
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	var d EnumInt = 5
	fmt.Println(a, b, c, d)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	// 二进制输出
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
	// String() 方法的 ChipType 在使用上和普通的常量没有区别。
	// 当这个类型需要显示为字符串时，Go语言会自动寻找 String() 方法并进行调用
	fmt.Printf("%s %d", GPU, GPU)
}
