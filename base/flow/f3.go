/*
3. Go语言for range（键值循环）
    for key, val := range coll {
        ...
    }
    a. 通过 for range 遍历的返回值有一定的规律：
        - 数组、切片、字符串返回索引和值。
        - map 返回键和值。
        - 通道（channel）只返回通道内的值。
    b. val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
*/

package main

import "fmt"

//  一般循环
func rangeBasic() {
	// 数组
	var ar [6]int = [6]int{1, 2, 3, 4, 5, 6}
	// ar := [6]int{1, 2, 3, 4, 5, 6}
	for i, v := range ar {
		fmt.Printf("i: %d, v: %d \n", i, v)
	}
	fmt.Println("-------------")
	// 切片
	var sp []int = []int{0, 2, 1, 1}
	// sp := ar[:5]
	// sp := make([]int, 5)
	sp[1] = 5
	sp = append(sp, 9)
	for i, v := range sp {
		fmt.Printf("i: %d, v: %d \n", i, v)
	}
	fmt.Println("-------------")
	// map
	var mp map[string]int = map[string]int{"a": 1, "b": 2}
	mp1 := make(map[string]int, 6)
	mp1["a"] = 4
	mp1["b"] = 3
	for k, v := range mp {
		fmt.Printf("k: %s, v: %d \n", k, v)
	}
	for k, v := range mp1 {
		fmt.Printf("k: %s, v: %d \n", k, v)
	}
	fmt.Println("-------------")
	// 字符串
	var s string = "xxxxxxxxxxxx"
	ss := "yyyyyy"
	for i, v := range s {
		fmt.Printf("i: %d, v: %U \n", i, v)
	}
	for i, v := range ss {
		fmt.Printf("i: %d, v: %U \n", i, v)
	}
}

func rangeChannel() {
	fmt.Println("-------------")
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3

		close(c)
	}()
	for v := range c {
		fmt.Printf("v: %v \n", v)
	}
}

func rangeSpecial() {
	fmt.Println("-------------")
	var ints []int = []int{1, 2, 3, 4, 5}
	fmt.Println(ints)
	for i, v := range ints {
		// 对应索引的值拷贝，因此它一般只具有只读性质
		fmt.Println(v)

		ints[i] = 5
	}
	fmt.Println(ints)
}

func main() {
	rangeBasic()
	rangeSpecial()
	rangeChannel()
}
