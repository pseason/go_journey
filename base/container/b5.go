/*
初始化列表
	name := list.New()
	var name list.List

*/

package main

import (
	"container/list"
	"fmt"
)

// list 初始
func initList() {
	// 声明 var name list.List
	var lis list.List
	lis.PushFront(1)
	// 声明 name := list.New()
	liis := list.New()
	liis.PushFront(2)

	fmt.Println(lis, liis)
}

// 列表操作
func listOpt() {
	// 初始化
	var liss list.List
	lis := list.New()
	// 添加元素
	liss.PushFront(9)
	liss.PushFront(10)

	lis.PushFront(1)
	lis.PushFront(2)
	lis.PushBack(3)
	// 循环list
	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 向 尾部、首部添加 list
	lis.PushBackList(&liss)
	lis.PushFrontList(&liss)

	// 循环list
	fmt.Println("------------------")
	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

// list 移除
func listRemove() {
	// 方式一
	var li list.List
	li.PushFront(1)
	li.PushFront(2)
	for l := li.Front(); l != nil; l = l.Next() {
		fmt.Printf("value: %v \n", l.Value)
	}
	fmt.Printf("init li len: %d \n", li.Len())
	for li.Len() > 0 {
		li.Remove(li.Front())
	}
	fmt.Printf("remove after li len: %d \n", li.Len())
	// 方式二
	lii := list.New()
	lii.PushFront(4)
	lii.PushFront(5)
	for l := lii.Front(); l != nil; l = l.Next() {
		fmt.Printf("value: %v \n", l.Value)
	}
	fmt.Printf("init lii len: %d \n", lii.Len())
	var next *list.Element
	for l := lii.Front(); l != nil; l = next {
		next = l.Next()
		lii.Remove(l)
	}
	fmt.Printf("remove after lii len: %d \n", lii.Len())
}

func main() {
	initList()
	listOpt()
	listRemove()
}
