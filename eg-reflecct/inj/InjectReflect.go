package inj

import (
	"fmt"

	"github.com/codegangsta/inject"
)

func s1() {
	inj := inject.New()
	//实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	//函数反转调用
	value, _ := inj.Invoke(Format)
	println("return: ", value[0].String())
}

func s2() {
	//创建被注入实例
	s := Staff{}
	//控制实例的创建
	inj := inject.New()
	//初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	//实现对 struct 注入
	inj.Apply(&s)
	//打印结果
	fmt.Printf("s ＝ %v\n", s)
}

func Inject() {
	s1()
	fmt.Println("------------")
	s2()
	fmt.Println("------------")

}
