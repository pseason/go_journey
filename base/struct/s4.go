/*
Go语言使用事件系统实现事件的响应和处理
*/
package main

import "fmt"

// 事件注册map
var EventMap = make(map[string][]func(interface{}))

// 注册事件
func RegisterEvent(name string, event func(interface{})) {
	// 通过名字查找事件列表
	splice := EventMap[name]
	// 添加事件
	splice = append(splice, event)
	// 重新赋值
	EventMap[name] = splice
}

// 执行事件
func ExecEvent(name string, param interface{}) {
	// 通过名字查找事件列表
	splice := EventMap[name]
	// 循环执行事件
	for _, event := range splice {
		// 执行
		event(param)
	}
}

// 声明角色的结构体
type Actor struct {
}

// actor结构方法
func (a *Actor) onEvent(param interface{}) {
	fmt.Println("Actor onEvent: ", param)
}

// 方法
func globalEvent(param interface{}) {
	fmt.Println("globalEvent: ", param)
}

func main() {
	// 初始actor
	actor := new(Actor)
	// 注册事件
	RegisterEvent("onSkill", actor.onEvent)
	RegisterEvent("onSkill", globalEvent)
	// 执行事件
	ExecEvent("onSkill", "execute")
}
