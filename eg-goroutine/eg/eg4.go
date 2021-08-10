package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 原子函数

var (
	count1 int64
	wg     sync.WaitGroup
)

func eg4(d int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		atomic.AddInt64(&count1, 1) //安全的对counter加1

		fmt.Println("eg:", d, i, count1)

		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)
	go eg4(1)
	go eg4(2)

	wg.Wait()
	fmt.Println(count1)

}
