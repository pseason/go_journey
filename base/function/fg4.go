/*
go 匿名函数
*/

package main

import (
	"flag"
	"fmt"
)

func simple() {
	// 匿名函数，55为直接调用
	func(i int) {
		fmt.Println("i:", i)
	}(55)

	// 匿名函数
	f := func(i int) int {
		return i + 6
	}
	v := f(6)
	fmt.Println("v:", v)
}

// 默认参数
var skillParam = flag.String("skill", "perform", "--skill=perform")

func main() {
	simple()
	// 解析
	flag.Parse()
	//
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
		"perform": func() {
			fmt.Println("perform fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
