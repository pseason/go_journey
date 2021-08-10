package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 16:20
version 1.0.0
desc:
	基础实例
*/

var wg sync.WaitGroup

func work() {
	for {
		fmt.Println("work")
		time.Sleep(time.Second)
	}
	// 如何停止
	wg.Done()
}

func main() {
	wg.Add(1)
	// 开启线程
	go work()

	wg.Wait()

	fmt.Println("done")
}
