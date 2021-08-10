/*
b struct
*/

package b

import (
	"fmt"

	"springmars.com/factory/base"
)

type B struct {
}

func (b *B) Do() {
	fmt.Println("B Do something")
}

// 包匿名函数 (import _ "b")会被使用
func init() {
	base.Register("B", func() base.Class {
		return new(B)
	})
}
