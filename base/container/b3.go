/*
Go语言map（Go语言映射）
	a. map 声明
		1. var name map[keyType] valueType
		2. 使用make()创建
			map := make(map[keyType]valueType)
            map := make(map[keyType]valueType,cap) cap作为扩容参数
	b. map 删除元素、清空
        删除元素：delete(map,key)
        清空数据：使用make()创建新map
	c.
	d.
	e.
*/
package main

import "fmt"

// map 声明
func ms() {
	// 简单声明
	var ms1 map[string]int
	fmt.Println(ms1, len(ms1))
	// 声明并初始化
	var ms2 map[string]int = map[string]int{"a": 1, "b": 2}
	fmt.Println(ms2, len(ms2))
	ms2["c"] = 3
	ms2["d"] = 4
	fmt.Println(ms2, len(ms2))
	// 使用 := 符号创建
	ms := map[int]string{}
	ms[1] = "a"
	fmt.Println(ms, len(ms))
	// 使用make() 创建
	ms3 := make(map[string]int)
	fmt.Println(ms3, len(ms3))
	ms3["c"] = 3
	ms3["d"] = 4
	fmt.Println(ms3, len(ms3))
	// 循环map
	for k, v := range ms2 {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}
	// 切片或者数据作为map的value，作为一对多
	ms4 := make(map[int][]int)
	arr := []int{1, 2}
	ms4[1] = arr
	fmt.Println(ms4, len(ms4))
	ms5 := make(map[int]*[]int)
	ms5[1] = &arr
	fmt.Println(ms5, len(ms5), *ms5[1])
}

// map 删除 清空
func msd() {
	// 初始化map
	ms6 := map[int]string{}
	ms6[1] = "a"
	ms6[2] = "b"
	ms6[3] = "c"
	ms6[4] = "d"
	fmt.Println(ms6, len(ms6))
	// 使用delete()删除元素
	delete(ms6, 1)
	fmt.Println(ms6, len(ms6))
	// 全部删除 or make()新函数
	for k := range ms6 {
		delete(ms6, k)
	}
	ms6 = make(map[int]string)
}

func main() {
	ms()
	msd()
}
