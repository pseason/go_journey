/*
go 语言中的hash函数
*/
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func some() {
	TestString := "Hi, pandaman!"
	// md5
	md5Inst := md5.New()
	md5Inst.Write([]byte(TestString))
	res := md5Inst.Sum([]byte(""))
	fmt.Printf("%x \n", res)
	// sha
	shaInst := sha1.New()
	shaInst.Write([]byte(TestString))
	res = shaInst.Sum([]byte(""))
	fmt.Printf("%x \n", res)
}

func main() {
	some()
}
