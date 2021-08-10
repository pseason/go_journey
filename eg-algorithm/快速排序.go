package main

import (
	"fmt"
	"math/rand"
)

// 快速排序
func main() {
	array := make([]int, 0)
	fmt.Println(array)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000)
	}
	fmt.Println("快速排序前", array)
	array = quicksort(array)
	fmt.Println("快速排序后", array)
}

// 快速排序
func quicksort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	first := array[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, first)
	for i := 1; i < len(array); i++ {
		if array[i] > first {
			high = append(high, array[i])
		} else if array[i] < first {
			low = append(low, array[i])
		} else {
			mid = append(mid, array[i])
		}
	}
	low, high = quicksort(low), quicksort(high)

	return append(append(low, mid...), high...)
}
