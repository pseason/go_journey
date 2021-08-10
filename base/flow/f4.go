/*
Go语言switch case语句

Go语言goto语句

*/

package main

import (
	"fmt"
	"strings"
)

//  switch基本语法
func switchBasic() {
	a := "a"
	switch a {
	case "a":
		fmt.Println("value is a")
	case "c", "b":
		fmt.Println("value is c or b")
	}

	i := 10
	switch {
	case i < 9:
		fmt.Println("value is less then 9")
	case i > 9 && i < 100:
		fmt.Println("value is greater then 9 and less then 100")
	case i > 100:
		fmt.Println("value is greater then 100")
	}

	si := "abc"
	switch {
	case strings.Contains(si, "a"):
		fmt.Println("value is contain a")
		fallthrough
	case strings.Contains(si, "d"):
		fmt.Println("value is contain d")
	case strings.Contains(si, "e"):
		fmt.Println("value is contain e")
	}

}

// goto 基础用法
func gotoBasic() {
	sum := 0
	for {
		sum++
		if sum > 6 {
			goto BLable
		}
	}
BLable:
	fmt.Println("goto after sum value is ", sum)

	if strings.Contains("abcdefgh", "cd") {
		goto IBLable
	}
IBLable:
	fmt.Println("value is contain cd")

}

func main() {

	switchBasic()

	gotoBasic()

}
