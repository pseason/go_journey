package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 16:30
version 1.0.0
desc:
	通道方式
*/

var wg sync.WaitGroup

var c = make(chan int)

func work() {
	for {
		fmt.Println("work")
		time.Sleep(time.Second)
		select {
		case <-c:
			wg.Done()
		default:

		}
	}
}

func main() {
	wg.Add(1)

	go work()

	time.Sleep(time.Second * 3)

	c <- 1

	close(c)

	wg.Wait()

	fmt.Println("done")
}
