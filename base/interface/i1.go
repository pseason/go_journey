/*
Go语言实现接口的条件
        - 接口的方法与实现接口的类型方法格式一致
		- 接口中所有方法均被实现
*/
package main

import "fmt"

type Cat interface {
	Miao(name string)
	Jump(hight int)
}

type BlackCat struct {
	Age  int
	Name string
}

type WhiteCat struct {
	Age  int
	Name string
}

func (b *BlackCat) Miao(name string) {
	fmt.Println(name, "BlackCat miao miao .....")
}

func (b *BlackCat) Jump(hight int) {
	fmt.Println(hight, "BlackCat jump .....")
}

func (b *WhiteCat) Miao(name string) {
	fmt.Println(name, "WhiteCat miao miao .....")
}

func main() {
	var cat Cat

	bc := new(BlackCat)
	wc := new(WhiteCat)

	cat = bc
	cat.Miao("lili")
	cat.Jump(12)
	// WhiteCat 未实现 Cat.Jump 方法，编译不通过
	cat = wc
	cat.Miao("liumei")

}
