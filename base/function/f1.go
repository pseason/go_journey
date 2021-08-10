/*
Go 语言函数（func）

*/

package main

import "fmt"

// 形式参数
func xx(x, y int) {
	fmt.Println(x, y)
}

// 形式参数 + 返回值
func xxx(x, y int) int {
	fmt.Println(x, y)
	return x + y
}

// 形式参数 + 返回值： 指定z接受
func xxxx(x, y int) (z int) {
	z = x + y
	return
}

// 形式参数（空白标识符_可以强调某个参数未被使用） + 返回值
func xxxxx(x int, _ int) int {
	return x + 10
}

func main() {
	xx(5, 6)
	fmt.Println(xxx(5, 8))
	c := xxxx(5, 8)
	fmt.Println(c)
	fmt.Println(xxxx(5, 8))

}
