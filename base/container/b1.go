/*
 数组（slice 切片）
a. 数组的声明语法如下：
        var 数组变量名 [元素数量]Type
        语法说明如下所示：
            - 数组变量名：数组声明及使用时的变量名。
            - 元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，
            元素数量不能含有到运行时才能确认大小的数值。
            - Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组
b. 多维数组
	声明多维数组的语法如下所示：
		var array_name [size1][size2]...[sizen] array_type
	其中，array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度
*/
package main

import "fmt"

// 数组
func ba() {
	var ints [5]int
	ints[2] = 15
	fmt.Println(ints)
	// 循环
	for i, v := range ints {
		fmt.Printf("index: %d, value: %v \n", i, v)
	}
	fmt.Println("--------------------------")
	for i := 0; i < len(ints); i++ {
		fmt.Printf("index: %d, value: %v \n", i, ints[i])
	}
	// 直接赋值定义
	var q [3]int = [3]int{1, 2, 3}
	var w [3]int = [3]int{1, 2}
	fmt.Printf("q[3]: %v \n", q[2])
	fmt.Printf("w[2]: %v \n", w[2])
	// 如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	e := [...]int{1, 2, 4, 5}
	fmt.Println(e)
	// 比较两个数组是否相等
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	// "true false false"
	fmt.Println(a == b, a == c, b == c)
}

// 多维数组
func bam() {
	// 直接声明
	var q1 [4][2]int
	q1[2][1] = 5
	q1[1] = [2]int{4, 2}
	fmt.Println(q1)
	// 指定位置声明
	var q2 [4][2]int
	q2 = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	q3 := [4][2]int{1: {1, 2}, 3: {3, 4}}
	// 声明并初始化数组中指定的元素
	q4 := [4][2]int{1: {1: 55}, 2: {0: 69}}
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)
	for i, v := range q4 {
		for j, v1 := range v {
			fmt.Printf("array[%v][%v]: %v \n", i, j, v1)
		}
	}
}

func main() {
	ba()
	fmt.Println("--------------------------")
	// 多维
	bam()
}
