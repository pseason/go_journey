package main

import "fmt"

func cg5() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for data := range ch {
		fmt.Println(data)
		if data > 1 {
			break
		}
	}
}

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 非阻塞接收
	data, ok := <-ch
	fmt.Println(data, ok)

	fmt.Println("all done")

	cg5()
}
