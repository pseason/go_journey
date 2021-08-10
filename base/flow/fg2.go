/*
冒泡排序
二分查找
*/

package main

import "fmt"

/*
冒泡排序
*/
func bubblingSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			t := array[i]
			if array[i] > array[j] {
				array[i] = array[j]
				array[j] = t
			}
		}
		fmt.Println("bubbling sort round: ", array)
	}
}

/*
二分查找
*/
func binarySearch(array *[]int, leftIndex int, rightIndex int, search int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到目标数据")
		return
	}
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*array)[middle] > search {
		binarySearch(array, leftIndex, middle-1, search)
	} else if (*array)[middle] < search {
		binarySearch(array, middle+1, rightIndex, search)
	} else {
		fmt.Println("找到value位置: ", middle)
	}
}

func main() {
	// 冒泡
	array := [...]int{454, 45, 78, 35, 459, 74, 69, 58, 124, 191, 22, 79, 59, 100}
	fmt.Println("bubbling sort before:", array)
	bubblingSort(array[:])
	fmt.Println("bubbling sort after:", array)
	// 二分
	splice := array[:]
	binarySearch(&splice, 0, len(array)-1, 58)

}
