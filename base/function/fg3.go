/*
go方法的链式调用
*/

package main

import (
	"fmt"
	"strings"
)

func stringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		// 遍历每一个处理链
		for _, proc := range chain {
			str = proc(str)
		}
		// 将结果放回切片
		list[index] = str
	}
}

func removePrifix(str string) string {
	return strings.TrimPrefix(str, "a")
}

func main() {
	list := []string{
		"a xxx",
		"a yyy z",
		"a z zz",
	}

	// 构造方法 切片
	chain := []func(string) string{
		removePrifix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	stringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}

}
