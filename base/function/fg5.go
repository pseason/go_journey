/*
函数实现接口
*/

package main

import "fmt"

type SimpleInterface interface {
	Call(interface{})
}

type SimpleStruct struct {
}

type SimpleFunc func(interface{})

func (s *SimpleStruct) Call(p interface{}) {
	fmt.Println("from simple struct:", p)
}

func (sf SimpleFunc) Call(p interface{}) {
	sf(p)
}

func main() {
	var invoker SimpleInterface

	s := new(SimpleStruct)

	invoker = s

	invoker.Call("struct")

	invoker = SimpleFunc(func(v interface{}) {
		fmt.Println("from simple func:", v)
	})

	invoker.Call("func")

}
