/*
Go语言使用空接口实现可以保存任意值的字典
*/
package main

import "fmt"

// 数据结构
type Pair struct {
	// map类型 key,value 都为interface
	mps map[interface{}]interface{}
}

// set key
func (p *Pair) Set(k, v interface{}) {
	p.mps[k] = v
}

// get key
func (p *Pair) Get(k interface{}) interface{} {
	return p.mps[k]
}

// delete key
func (p *Pair) Delete(k interface{}) {
	delete(p.mps, k)
}

// clear
func (p *Pair) Clear() {
	p.mps = make(map[interface{}]interface{})
}

// visit
func (p *Pair) Visit(callback func(k, v interface{}) bool) {
	if callback == nil || p.mps == nil {
		return
	}
	for k, v := range p.mps {
		if !callback(k, v) {
			return
		}
	}

}

func NewPair() *Pair {
	p := new(Pair)
	p.mps = make(map[interface{}]interface{})
	return p
}

func main() {
	p := NewPair()
	// 添加游戏数据
	p.Set("My Factory", 60)
	p.Set("Terra Craft", 36)
	p.Set("Don't Hungry", 24)
	// 获取值及打印值
	favorite := p.Get("Terra Craft")
	fmt.Println("favorite:", favorite)
	// 遍历所有的字典元素
	p.Visit(func(key, value interface{}) bool {
		// 将值转为int类型，并判断是否大于40
		if value.(int) > 40 {
			// 输出很贵
			fmt.Println(key, "is expensive")
			return true
		}
		// 默认都是输出很便宜
		fmt.Println(key, "is cheap")
		return true
	})
}
