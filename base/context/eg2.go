package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 16:26
version 1.0.0
desc:
	全局变量方式
*/

var wg sync.WaitGroup
var exit bool

func work() {
	for {
		fmt.Println("work")
		time.Sleep(time.Second)

		if exit {
			break
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	go work()

	time.Sleep(time.Second * 3)
	exit = true

	wg.Wait()

	fmt.Println("done")
}
