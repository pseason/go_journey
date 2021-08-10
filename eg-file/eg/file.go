package eg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func encoderjsonfile() {

	ws := []Website{
		{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}},
		{"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}},
	}

	file, err := os.Create("website.json")
	if err != nil {
		fmt.Println("创建website.json错误", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(ws)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func decoderjsonfile() {
	file, err := os.Open("website.json")
	if err != nil {
		fmt.Println("打开website.json错误", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	// 读取
	var info []Website
	err = decoder.Decode(&info)
	if err == nil {
		fmt.Println(info)
	} else {
		fmt.Println("解码失败")
	}
}

func init() {
	encoderjsonfile()
	fmt.Println("-------------")
	decoderjsonfile()
	fmt.Println("-------------")

}
