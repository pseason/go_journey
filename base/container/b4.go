/*
Go语言sync.Map
	a. 声明sync.Map变量
		var xxx sync.Map
	sync.Map 有以下特性：
		- 无须初始化，直接声明即可。
		- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，
			Store 表示存储，Load 表示获取，Delete 表示删除。
		- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
			Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false

*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Map的相关操作
func syncms() {
	// 声明
	var smps sync.Map
	// 存储
	smps.Store(1, "a")
	smps.Store(2, "b")
	smps.Store(3, "c")
	fmt.Println(smps)
	// 获取
	v, ok := smps.Load(1)
	fmt.Println(v, ok)
	// 删除
	smps.Delete(3)
	fmt.Println(smps)
	// 循环
	smps.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//查询并且删除
	smps.Store(3, "c")
	pre, ok := smps.LoadAndDelete(3)
	fmt.Println(pre, ok)
	pre, ok = smps.LoadAndDelete(3)
	fmt.Println(pre, ok)
	// 查询并保存
	pre, ok = smps.LoadOrStore(3, "c")
	fmt.Println(pre, ok)
	aft, okk := smps.Load(3)
	fmt.Println(aft, okk)
	pre, ok = smps.LoadOrStore(3, "d")
	fmt.Println(pre, ok)
	aft, okk = smps.Load(3)
	fmt.Println(aft, okk)
}

// 声明
var smp sync.Map

//  sync test
func syncTest1() {
	for i := 0; i < 50; i++ {
		fmt.Println("11111111111")
		v, _ := smp.Load(1)
		smp.Store(1, v.(int)+1)
	}
}
func syncTest2() {
	for i := 0; i < 50; i++ {
		fmt.Println("22222222222")
		v, _ := smp.Load(1)
		smp.Store(1, v.(int)+1)
	}
}
func syncTest3() {
	for i := 0; i < 50; i++ {
		fmt.Println("33333333333")
		v, _ := smp.Load(1)
		smp.Store(1, v.(int)+1)
	}
}

func main() {
	syncms()
	// 存储
	smp.Store(1, 1)
	smp.Store(2, 1)
	smp.Store(3, 1)

	go syncTest1()
	go syncTest2()
	go syncTest3()

	time.Sleep(time.Second * 3)
	v, ok := smp.Load(1)
	fmt.Println(v, ok)
}
