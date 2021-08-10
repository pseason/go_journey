package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 17:11
version 1.0.0
desc:

*/

var wg sync.WaitGroup

var ctx, cancel = context.WithCancel(context.Background())

func work() {

	go work1()

	for {
		fmt.Println("work")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			wg.Done()
		default:

		}
	}
}

func work1() {
loop:
	for {
		fmt.Println("work1")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break loop
		default:

		}
	}
	fmt.Println("work1 over")
}

func main() {
	wg.Add(1)

	go work()

	time.Sleep(time.Second * 3)

	cancel()

	wg.Wait()

	fmt.Println("done")
}
