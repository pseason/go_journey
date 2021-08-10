package test

import "testing"

func TestPlus(t *testing.T) {
	c := Plus(4, 5)
	if c != 9 {
		t.Error("测试 error")
	}
}
