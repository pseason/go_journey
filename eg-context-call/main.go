package main

import (
	"fmt"
	"math/rand"
	"springmars.com/context-call/call"
	"time"
)

/*
@author pengshuo
@date 2021/8/6 19:32
version 1.0.0
desc:

*/

func main() {
	call.Init()

	for i := 0; i < 10; i++ {
		go func(value int) {
			id := call.AfterFutureTask(func(data interface{}) {
				fmt.Println("task receive:", data.(int))
			}, value)

			fmt.Println("put future: ", value, id)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))

			call.Call(id)
		}(i)
	}

	time.Sleep(time.Second * 5)
}
