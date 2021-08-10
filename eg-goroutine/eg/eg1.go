package main

import (
	"fmt"
	"time"
)

func round() {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		// 延迟1秒
		time.Sleep(time.Second)
	}
}

func main() {

	go round()

	var input string

	fmt.Scanln(&input)
	fmt.Println(input)
}
