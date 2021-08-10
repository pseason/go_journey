package main

import "fmt"

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	fmt.Printf("t 变量的地址是: %x\n", &t)

	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t

	fmt.Printf("%p, %v \n", &t, t)
}

func xx() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func main() {

	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	fmt.Printf("%p, %v \n", &x, x)
	fmt.Printf("%p, %v \n", &y, y)
	fmt.Println("-----------------")
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println("-----------------")
	fmt.Printf("%p, %v \n", &x, x)
	fmt.Printf("%p, %v \n", &y, y)
	fmt.Println("-----------------")

	xx()

}
