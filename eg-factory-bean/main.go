package main

import (
	"fmt"

	_ "springmars.com/factory/a"
	_ "springmars.com/factory/b"
	"springmars.com/factory/base"
)

func main() {
	fmt.Println("---------")

	a := base.Create("A")
	a.Do()

	b := base.Create("B")
	b.Do()

}
