/*
go语言 defer
*/
package main

import (
	"fmt"
	"sync"
)

//  defer 先进后出
func simple() {
	fmt.Println("defer begin")
	defer fmt.Println("defer end")
	defer fmt.Println("-------")
}

// 关闭资源
var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

func readValue(key string) int {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func f1() {
	fmt.Println("f1------")
	defer fmt.Println("f1 defer-----------")
	f2()
}
func f2() {
	fmt.Println("f2------")
	defer fmt.Println("f2 defer-----------")
}

func main() {
	simple()

	f1()
}
