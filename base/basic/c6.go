package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
字符串及其操作
	定义
	字符串拼接 “+”
	定义多行字符串 ``
	字符串长度
		len()、utf8.RuneCountInString
	字符串遍历
	字符串截取
		字符串索引比较常用的有如下几种方法：
		strings.Index：正向搜索子字符串。
		strings.LastIndex：反向搜索子字符串。
		搜索的起始位置可以通过切片偏移制作。
	字符串修改
		修改字符串时，可以将字符串转换为 []byte 进行修改。
		[]byte 和 string 可以通过强制类型转换互转
	字符串拼接
        使用类java StringBuffer, bytes.Buffer
	base64 编解码
*/

func main() {
	// 定义
	var str string = "hello go 学习进展顺利"
	fmt.Println(str)
	// 字符串拼接 “+”
	fmt.Println(str + " \u6253")
	// 定义多行字符串 ``
	var line string = `go
	java
	python
	`
	fmt.Println(line)
	// 字符串长度计算
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	// 字符串遍历
	for _, s := range str {
		fmt.Printf("%c --> %d \n", s, s)
	}
	// 字符串截取
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])
	// 字符串修改
	var cm string = "hello xxxx world"
	cmb := []byte(cm)
	for i := 5; i < 10; i++ {
		cmb[i] = ' '
	}
	fmt.Println(string(cmb))
	// 字符串拼接
	var sb bytes.Buffer
	sb.WriteString(cm)
	sb.WriteString(" go study")
	fmt.Println(sb.String())
	// base64 编解码
	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"
	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出
	fmt.Println(encodedMessage)
	// 解码
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}
