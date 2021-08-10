package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

var lock sync.Mutex

func counter() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("counter count:", count)
}

func main() {

	for i := 0; i < 10; i++ {
		go counter()
	}

	for {
		lock.Lock()
		c := count
		fmt.Println("main count:", count)
		lock.Unlock()
		// 允许其他goroutine执行,执行完毕，在执行此处
		runtime.Gosched()

		if c > 9 {
			break
		}
	}

}
