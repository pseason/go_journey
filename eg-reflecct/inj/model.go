package inj

import (
	"fmt"
)

type S1 interface{}
type S2 interface{}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func Format(name string, company S1, level S2, age int) string {
	return fmt.Sprintf("name ＝ %s, company=%s, level=%s, age ＝ %d", name, company, level, age)
}
