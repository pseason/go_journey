package pkgc

import "b/pkgb"

func Add(a, b int) int {
	return pkgb.ExcuteAdd(1, 3)
}
