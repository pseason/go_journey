package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := make([]int, 10)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println("排序前:", array)
	bubbling(array)
	fmt.Println("排序后:", array)
}

// 冒泡排序
func bubbling(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
		fmt.Println("排序中:", array)
	}
}

// 冒泡排序 去最大值
func bubbleMax(array []int) int {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
		}
	}
	return array[len(array)-1]
}
