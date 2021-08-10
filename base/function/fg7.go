/*
go 通过内存缓存来提升性能
*/

package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// 普通计数
func countTime() {
	result := 0
	start := time.Now()
	for i := 0; i < 41; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	delta := time.Since(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}

// 内存计数
const LIM = 41

var fibs [LIM]uint64

func fibonacciLim(n int) uint64 {
	if fibs[n] != 0 {
		return fibs[n]
	}
	if n <= 2 {
		fibs[n] = 1
	} else {
		fibs[n] = fibonacciLim(n-1) + fibonacciLim(n-2)
	}
	return fibs[n]
}

func countLimTime() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacciLim(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	delta := time.Since(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}

func main() {
	countTime()
	countLimTime()
}
