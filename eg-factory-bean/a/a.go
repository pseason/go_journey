/*
a struct
*/
package a

import (
	"fmt"

	"springmars.com/factory/base"
)

type A struct {
}

func (a *A) Do() {
	fmt.Println("A Do something")
}

func init() {
	base.Register("A", func() base.Class {
		return new(A)
	})
}
