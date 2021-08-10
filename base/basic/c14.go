/*
Go语言strconv包：字符串和数值类型的相互转换
        1. Itoa()：整型转字符串
        2. Atoi()：字符串转整型
        3. Parse 系列函数
            ParseBool()
            ParseInt()
            ParseUnit()
            ParseFloat()
        4. Format 系列函数
            FormatBool()
            FormatInt()
            FormatUint()
            FormatFloat()
        5. Append 系列函数
            AppendBool()
            AppendFloat()
            AppendInt()
            AppendUint()
*/
package main

import (
	"fmt"
	"strconv"
)

func str4int() {
	a := 9
	b := "898"
	c := strconv.Itoa(a)
	fmt.Printf("%T: %s \n", c, c)
	d, err := strconv.Atoi(b)
	if err == nil {
		fmt.Printf("%T: %#v \n", d, d)
	}
}

func parse() {
	// parse bool 只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	a := "1"
	b, err := strconv.ParseBool(a)
	if err == nil {
		fmt.Printf("%T: %#v \n", b, b)
	}
	c := "xxx"
	d, err := strconv.ParseBool(c)
	if err != nil {
		fmt.Printf("%v \n", err)
	} else {
		fmt.Printf("%T: %#v \n", d, d)
	}
	// parse int
	// ParseInt() 函数用于返回字符串表示的整数值（可以包含正负号），函数签名如下：
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// 参数说明：
	// base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
	// bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
	// 返回的 err 是 *NumErr 类型的，如果语法有误，err.Error = ErrSyntax，如果结果超出类型范围 err.Error = ErrRange。
	e := "-15"
	f, err := strconv.ParseInt(e, 10, 0)
	if err == nil {
		fmt.Printf("%T: %#v \n", f, f)
	} else {
		fmt.Printf("%v \n", err)
	}
	// ParseUint() 函数的功能类似于 ParseInt() 函数，但 ParseUint() 函数不接受正负号，用于无符号整型
	str := "11"
	num, err := strconv.ParseUint(str, 10, 0)
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}
	// ParseFloat()  函数用于将一个表示浮点数的字符串转换为 float 类型
	// 参数说明：
	// 如果 s 合乎语法规则，函数会返回最为接近 s 表示值的一个浮点数（使用 IEEE754 规范舍入）。
	// bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；
	// 返回值 err 是 *NumErr 类型的，如果语法有误 err.Error=ErrSyntax，如果返回值超出表示范围，返回值 f 为 ±Inf，err.Error= ErrRange
	str1 := "3.1415926"
	num1, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num1)
	}
}

func format() {
	s1 := strconv.FormatBool(true)
	// 参数说明：
	// bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入。
	// fmt 表示格式，可以设置为“f”表示 -ddd.dddd、“b”表示 -ddddp±ddd，指数为二进制、“e”表示 -d.dddde±dd 十进制指数、“E”表示 -d.ddddE±dd 十进制指数、“g”表示指数很大时用“e”格式，否则“f”格式、“G”表示指数很大时用“E”格式，否则“f”格式。
	// prec 控制精度（排除指数部分）：当参数 fmt 为“f”、“e”、“E”时，它表示小数点后的数字个数；当参数 fmt 为“g”、“G”时，它控制总的数字个数。如果 prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f。
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// 参数 i 必须是 int64 类型，参数 base 必须在 2 到 36 之间
	s3 := strconv.FormatInt(-42, 16)
	// 与 FormatInt() 函数的功能类似，但是参数 i 必须是无符号的 uint64 类型
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s1, s2, s3, s4)
}

func append() {
	// 声明一个slice
	b10 := []byte("int (base 10):")

	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

func main() {
	// string with number
	str4int()
	// parse
	parse()
	// format
	format()
	// append
	append()
}
