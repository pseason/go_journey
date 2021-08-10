package main

import "fmt"

/**
指针类型
	Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作，当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值
	数据取内存地址：ptr x := &y
	内存地址转取值：z := *x
*/
func main() {
	var str string = "i love you"
	prt := &str
	fmt.Printf("prt type: %T ,address: %p \n", prt, prt)

	va := *prt
	fmt.Printf("va type: %T, val: %s", va, va)

}
