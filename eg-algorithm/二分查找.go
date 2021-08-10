package main

import (
	"fmt"
	"math/rand"
)

// 二分查找
func main() {
	array := make([]int, 10)
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(10000)
	}
	fmt.Println("冒泡排序前:", array)
	bubblesort(array)
	fmt.Println("冒泡排序后:", array)
	// 二分查找
	fd := array[rand.Intn(10)]
	rd := binarysearch(array, fd)
	fmt.Println("二分查找数据：", fd, rd)
}

// 冒泡排序
func bubblesort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

// 二分查找
func binarysearch(array []int, fvalue int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] > fvalue {
			high = mid - 1
		} else if array[mid] < fvalue {
			low = mid + 1
		} else {
			return array[mid]
		}
	}
	return -1
}
