/*
 类与接口的关系
        - 一个类可以实现多个接口
        - 多个类型可以实现相同接口
*/
package main

import "fmt"

type Worker interface {
	Codeing()
	Checking()
}

type SmallTeacher struct {
}

type BigTeacher struct {
	SmallTeacher
}

func (bt *BigTeacher) Codeing() {
	fmt.Println("BigTeacher Codeing")
}

func (st *SmallTeacher) Checking() {
	fmt.Println("SmallTeacher Checking")
}

// 利用嵌套类实现多接口
func main() {
	var worker Worker
	bt := new(BigTeacher)
	worker = bt

	worker.Checking()
	worker.Codeing()
}
