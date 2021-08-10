/*
go 结构体
*/
package main

import "fmt"

// 定义结构体
type Student struct {
	Id   int32
	Name string
	Age  int
}

//  构造函数
func CreateNewStudent(id int32, name string, age int) *Student {
	return &Student{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

// 成员变量set属性方法，类似继承
func (s *Student) SetName(name string) {
	s.Name = name
}

// 初始化结构体
func initStudent() {
	fmt.Println(nil == new(Student))
	// 使用new关键字初始化
	lili := new(Student)
	lili.Age = 18
	lili.SetName("lili")
	fmt.Println(lili)
	// 使用&取指针初始化
	lucy := &Student{Age: 18}
	lucy.Name = "lucy"
	fmt.Println(lucy)
	// 使用构造方法
	jack := CreateNewStudent(12, "jack", 19)
	fmt.Println(jack)
}

func main() {
	initStudent()
}
