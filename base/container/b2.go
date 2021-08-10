/*
切片（slice）
a.从数组生成切片
	slice [开始位置 : 结束位置]
	语法说明如下：
		slice：表示目标切片对象；
		开始位置：对应目标切片对象的索引；
		结束位置：对应目标切片的结束索引
b.直接声明新切片
	var name []type -> var ints []int
        其中 name 表示切片的变量名，Type 表示切片对应的元素类型
f. 使用 make() 函数构造切片
	如果需要动态地创建一个切片，可以使用 make() 内建函数
		make( []Type, size, cap )
	其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，
	这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题
g. 使用append()方法为切片添加元素
	slice = append(slice,value)

	切片默认cap为6，容量的扩展规律是按,长度少于1024时,
        容量的2倍数进行扩充,超过1024的，1.25倍扩容
h. 切片拷贝复制(copy)
	copy() 函数的使用格式如下
            copy( destSlice, srcSlice )
i. 切片删除
	1.从开头位置删除
		利用切片特性              a = a[N:]
		利用append()函数          append(a[:0], a[N:]...)
		利用copy()函数            a[:copy(a,a[N:])]
	2.从中间位置删除
		利用append()函数          append(a[:i],a[i+n:]...)
		利用copy()函数            a[:i+copy(a[:i],a[i+N:])]
	3.从尾部删除
		利用切片特性              a = a[:len(a) - N]
*/
package main

import (
	"fmt"
)

// 从数组生成切片
func bs() {
	// 定义数组
	var w1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	// 从指定范围中生成切片
	sp1 := w1[0:2]
	fmt.Println(sp1)
	// 缺省
	fmt.Println(w1[:5])
	fmt.Printf("%T %v \n", w1[1:], w1[1:])
	fmt.Printf("%T: %v \n", w1, w1) //数组class to string默认带长度
	// 表示原有的切片
	fmt.Println(w1[:])
	// 重置切片，清空拥有的元素
	sp1 = sp1[0:0]
	fmt.Println(sp1)
}

// 直接定义切片
func bsa() {
	// 声明整形切片
	var ints []int
	// 声明一个空切片
	var strs []string = []string{}
	fmt.Println(ints, strs)
	fmt.Printf("%T: %v \n", ints, ints) //切片class to string默认不带长度
	// 输出3个切片大小
	fmt.Println(len(ints), len(strs))
	// 切片判定空的结果
	fmt.Println(ints == nil, strs == nil)
}

// 使用make函数创建切片
func bsm() {
	spp1 := make([]int, 5)
	spp2 := make([]int, 5, 10)
	fmt.Println(spp1, spp2)
	fmt.Println(len(spp1), len(spp2))
	fmt.Printf("%T: %v \n", spp1, spp1)
	fmt.Printf("%T: %v \n", spp2, spp2)
}

// 使用append()方法为切片添加元素
func bsAppend() {
	var sparr [6]int = [6]int{11, 12, 13, 14, 15, 16}
	//定义切片
	var spsx []int = sparr[:1]
	fmt.Println(spsx, len(spsx), cap(spsx))
	var sps1 []int = sparr[:4]
	fmt.Println(sps1, len(sps1), cap(sps1))
	// 追加1个元素
	sps1 = append(sps1, 1)
	fmt.Println(sps1, len(sps1), cap(sps1))
	// 追加多个元素, 手写解包方式
	sps1 = append(sps1, 2, 2, 3, 4)
	fmt.Println(sps1, len(sps1), cap(sps1))
	// 追加一个切片, 切片需要解包
	sps1 = append(sps1, []int{5, 6, 7}...)
	fmt.Println(sps1, len(sps1), cap(sps1))
	sps1 = append(sps1, 1)
	fmt.Println(sps1, len(sps1), cap(sps1))
	sps1 = append(sps1, []int{5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 3, 2, 4, 5}...)
	fmt.Println(sps1, len(sps1), cap(sps1))
}

// 切片拷贝复制
func bsCopy() {
	// 定义切片
	var a []int = []int{1, 2, 3, 4, 5}
	var b []int = []int{6, 7, 8}
	var c []int = []int{6, 7, 8}
	// 只会复制slice1的前3个元素到slice2中
	copy(b, a)
	fmt.Println(a, b)
	// 只会复制slice2的3个元素到slice1的前3个位置
	copy(a, c)
	fmt.Println(a, c)
}

// 对切片的引用和复制操作后对切片元素的影响
func bsCopy2() {
	// 设置元素数量为1000
	const size = 12
	// 预分配足够多的元素切片
	a := make([]int, size)
	// 将切片赋值
	for i := 0; i < size; i++ {
		a[i] = i
	}
	// 引用切片数据
	ca := a
	// 预分配足够多的元素切片
	b := make([]int, size)
	// 将数据复制到新的切片空间中
	copy(b, a)
	// 修改原始数据的第一个元素
	a[0] = 999
	fmt.Println(a, ca, b)
	//复制原始数据从4到6(不包含)
	copy(b, a[4:6])
	fmt.Println(a, ca, b)
}

//切片元素删除
func bsDelete() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 从开头位置删除 b = b[N:] 删除开头N个元素
	b := make([]int, 8)
	copy(b, a)
	b = b[1:]
	fmt.Println(b)
	// 从开头位置删除 append(a[:0], a[N:]...) 删除开头N个元素
	b = make([]int, 8)
	copy(b, a)
	b = append(b[:0], b[1:]...)
	fmt.Println(b)
	// 从开头位置删除 a[:copy(a,a[N:])] 删除开头N个元素
	b = make([]int, 8)
	copy(b, a)
	b = b[:copy(b, b[1:])]
	fmt.Println(b)
	// 从中间位置删除 append(a[:n], a[n+N:]...) 删除中间N个元素
	b = make([]int, 8)
	copy(b, a)
	b = append(b[:3], b[3+3:]...)
	fmt.Println(b)
	// 从中间位置删除 a[:i+copy(a[:i],a[i+n:])] 删除中间N个元素
	b = make([]int, 8)
	copy(b, a)
	b = b[:3+copy(b[3:], b[3+3:])]
	fmt.Println(b)
	// 从尾部删除 a[:len(a) -n] 删除尾部N个元素
	b = make([]int, 8)
	copy(b, a)
	b = b[:len(b)-3]
	fmt.Println(b)
	// 切片循环
	g := []int{10, 20, 30, 40}
	for i, v := range g {
		fmt.Printf("index: %v , value: %v , value-addr: %X , ele-addr: %X \n", i, v, &v, &g[i])
	}
	for i := 0; i < len(g); i++ {
		fmt.Printf("index: %v , value: %v , value-addr: %X ，ele-addr: %X  \n", i, g[i], &g[i], &g[i])
	}
}

func main() {
	bs()
	bsa()
	bsm()
	bsAppend()
	bsCopy()
	bsCopy2()
	bsDelete()
}
