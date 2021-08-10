/*
map多键索引——多个数值条件可以同时查询
	2. 利用 map 特性的多键索引及查询
            使用结构体进行多键索引和查询比传统的写法更为简单，
            最主要的区别是无须准备哈希函数及相应的字段无须做哈希合并

*/
package main

import "fmt"

type Profile struct {
	Name string
	Age  int
}

func main() {
	list := []Profile{
		{Name: "xxx", Age: 18},
		{Name: "yyy", Age: 19},
	}

	maps := make(map[Profile]Profile)

	for _, v := range list {
		maps[v] = v
	}

	fmt.Println(maps[Profile{Name: "xxx", Age: 18}])

}
