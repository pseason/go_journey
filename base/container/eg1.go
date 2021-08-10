/*
map多键索引——多个数值条件可以同时查询
	1. 基于哈希值的多键索引及查询
            传统的数据索引过程是将输入的数据做特征值。这种特征值有几种常见做法：
                - 将特征使用某种算法转为整数，即哈希值，使用整型值做索引。
                将特征转为字符串，使用字符串做索引。
                - 数据都基于特征值构建好索引后，就可以进行查询。查询时，重复这个过程，
                将查询条件转为特征值，使用特征值进行查询得到结果

*/
package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

func (p *Profile) hash() int {
	return simpleHash(p.Name) + p.Age*1000000
}

func simpleHash(str string) (ret int) {
	// 遍历字符串中的每一个ASCII字符
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i]
		// 将字符的ASCII码相加
		ret += int(c)
	}
	return
}

// 全局map
var maps = make(map[int][]*Profile)

//存入 map[index][]profile
func buildIndex(list []*Profile) {
	for _, v := range list {
		// fmt.Printf("mem-addr: %p ,value: %v\n", v, *v)
		key := v.hash()
		spl := maps[key]
		spl = append(spl, v)
		maps[key] = spl
	}
}

// 查询数据
func queryData(name string, age int) *Profile {
	profile := Profile{Name: name, Age: age}
	key := profile.hash()
	spl := maps[key]
	for _, v := range spl {
		if name == v.Name && age == v.Age {
			return v
		}
	}
	return nil
}

func main() {
	// 创建 list
	list := []*Profile{
		{Name: "lili", Age: 18, Married: false},
		{Name: "lucy", Age: 28},
		{Name: "liumei", Age: 28},
	}
	// 创建索引
	buildIndex(list)
	// 查询
	p := queryData("liumei", 28)
	fmt.Printf("value: %+v", *p)
	// fmt.Printf("mem-addr: %p value: %v\n", p, *p)
}
