/*
处理
*/
package pkgb

import "b/pkga"

func ExcuteAdd(a, b int) int {
	return pkga.Add(a, b)
}

func ExcuteMinus(a, b int) int {
	return pkga.Minus(a, b)
}
