/*
go 锁
*/
package m

import (
	"fmt"
	"sync"
	"time"
)

func Add4Unlock() {
	a := 0
	for i := 0; i < 1000; i++ {
		go func(i int) {
			a++
			fmt.Println("a: ", a)
		}(i)
	}
	time.Sleep(time.Second)
}

func Add4Lock() {
	var lock sync.Mutex
	a := 0
	for i := 0; i < 1000; i++ {
		go func(i int) {
			// 加锁
			lock.Lock()
			defer lock.Unlock()

			a++
			fmt.Println("a: ", a)
		}(i)
	}
	time.Sleep(time.Second)
}

func Lock4Goroutine() {
	ch := make(chan struct{}, 2)
	var lock sync.Mutex
	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("goroutine2: 等待解锁")
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("goroutine2: 欧耶，也我解锁了")
		ch <- struct{}{}
	}()

	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}
