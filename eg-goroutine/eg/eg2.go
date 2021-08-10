package main

import (
	"fmt"
	"time"
)

func round2() {

	var times int

	// 匿名函数
	go func() {
		for {
			times++
			fmt.Println("tick:", times)
			// 延迟1秒
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	round2()

	var input string

	fmt.Scanln(&input)
	fmt.Println(input)
}
