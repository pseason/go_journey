/*
为任意类型添加方法
        type MyName Type:
            type MyInt int
            func (i MyInt) IsZero() bool{
                return i == 0
            }
*/
package main

import "fmt"

type MyInt int

func (i MyInt) IsZero() bool {
	return i == 0
}

func (i MyInt) Add(add int) int {
	return int(i) + add
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())

	b = MyInt(b.Add(5))

	fmt.Println(b)
}
