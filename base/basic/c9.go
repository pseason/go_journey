package main

import "fmt"

/**
变量逃逸分析
	1.栈（Stack）是一种拥有特殊规则的线性表数据结构
	2.栈只允许从线性表的同一端放入和取出数据，按照后进先出（LIFO，Last InFirst Out）的顺序
	3.栈可用于内存分配，栈的分配和回收速度非常快
	4.堆分配内存和栈分配内存相比，堆适合不可预知大小的内存分配。
	  但是为此付出的代价是分配速度较慢，而且会形成内存碎片
	*5.Go语言将这个过程整合到了编译器中，命名为“变量逃逸分析”。通过编译器分析代码的特征和代码的生命周期，
	  决定应该使用堆还是栈来进行内存分配
	6.编译器觉得变量应该分配在堆和栈上的原则是：
		变量是否被取地址；
		变量是否发生逃逸。
	go build -gcflags "-m -l" file.go
*/

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

// 空函数, 什么也不做
func void() {
}

type Data struct {
}

func point() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c
}

func main() {
	// 声明a变量并打印
	var a int
	// 调用void()函数
	void()
	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))
	// 打印指针
	fmt.Println(point())
}
