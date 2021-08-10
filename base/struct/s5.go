/*
go json 数据
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name string
	Age  int
}

func simple() {
	t := new(Teacher)
	t.Age = 32
	t.Name = "LiMei"
	// 生成json
	b, err := json.Marshal(t)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println("生成json error", err)
	}

	limei := new(Teacher)
	str := "{\"Name\":\"LiMei\",\"Age\":32"
	err = json.Unmarshal([]byte(str), limei)
	if err == nil {
		fmt.Println(limei)
	} else {
		fmt.Println("解析json error:", err)
	}

}

func main() {
	simple()
}
