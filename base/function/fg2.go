/*
 函数中的参数传递效果测试
*/

package main

import "fmt"

type Data struct {
	complex []int
	a       IntData
	b       *IntData
}

type IntData struct {
	a int
}

func simple(d Data) Data {
	fmt.Printf("value: %+v \n", d)
	fmt.Printf("point: %p \n", &d)
	d.a = IntData{
		7,
	}
	// 更改指针地址指向值的值
	*d.b = IntData{
		10,
	}
	return d
}

func complex(d *Data) *Data {
	fmt.Printf("value: %+v \n", d)
	fmt.Printf("point: %p \n", &d)
	d.a = IntData{
		6,
	}
	// 更改指针地址指向值的值
	*d.b = IntData{
		9,
	}
	return d
}

func main() {
	// 准备传入函数的结构
	in := Data{
		complex: []int{1, 2, 3},
		a: IntData{
			5,
		},
		b: &IntData{1},
	}
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in point: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := simple(in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out point: %p\n", &out)
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in point: %p\n", &in)
	fmt.Printf("in point: %v\n", *in.b)

	fmt.Println("------------------------------")

	in1 := Data{
		complex: []int{1, 2, 3},
		a: IntData{
			5,
		},
		b: &IntData{1},
	}
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in1)
	// 输入结构的指针地址
	fmt.Printf("in point: %p\n", &in1)
	// 传入结构体，返回同类型的结构体
	out1 := *complex(&in1)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out1)
	// 输出结构的指针地址
	fmt.Printf("out point: %p\n", &out1)
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in1)
	// 输入结构的指针地址
	fmt.Printf("in point: %p\n", &in1)

	fmt.Printf("in point: %v\n", *in1.b)

}
