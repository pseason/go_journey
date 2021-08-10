/*
带有父子关系的结构体的构造和初始化——模拟父级构造调用

*/
package main

import "fmt"

// 父类
type Cat struct {
	Age  int
	Name string
}

func (c *Cat) miaomiao() {
	fmt.Println(c.Name, "miao miao...")
}

// 子类
type BlackCat struct {
	Cat
	color string
}

func main() {
	bc := new(BlackCat)
	bc.Age = 16
	bc.Name = "xiaohua"
	bc.color = "black-white"
	bc.miaomiao()
	fmt.Println(bc)
}
