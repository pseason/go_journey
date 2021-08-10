/*
接口之间的相互转换
*/
package main

import "fmt"

func simple() {
	var i interface{} = 5
	var b int = i.(int)
	fmt.Println(i, b)

	var s interface{} = "xxxx"
	var ss string = s.(string)
	fmt.Println(s, ss)

}

func main() {
	simple()
}
