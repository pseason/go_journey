/*
1. go 语言if else（分支结构）
    a. 基本： if 条件 {} else 条件 {} else{}
    b. 特殊： if 执行语句; 条件 {} else 条件 {} else{}
*/

package main

import "fmt"

// 基本
func ifBasic() {
	var i int = 10
	// if 条件判断
	if i == 5 {
		fmt.Println("value is 5")
	} else if i == 10 {
		fmt.Println("value is 10")
	} else {
		fmt.Println("value is not here")
	}
}

// 特殊
func ifSpecial() {
	var ii int = 10
	// if 执行语句；条件判断
	if ii++; ii < 5 {
		fmt.Println("value is less then 5")
	} else if ii > 10 {
		fmt.Println("value is greater then 10")
	} else {
		fmt.Println("value is not here")
	}
}

func main() {
	ifBasic()
	ifSpecial()
}
