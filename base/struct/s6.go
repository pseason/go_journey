/*
go io
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func simple() {
	data := []byte("C语言中文网")
	br := bytes.NewReader(data)
	r := bufio.NewReader(br)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), err)
}

func simple1() {
	data := []byte("C语言中文网")
	br := bytes.NewReader(data)
	r := bufio.NewReader(br)
	b, err := r.ReadByte()
	fmt.Println(string(b), err)
}

func simple2() {
	data := []byte("C语言中文网,java语言中文网")
	br := bytes.NewReader(data)
	r := bufio.NewReader(br)
	b, err := r.ReadBytes(',')
	fmt.Println(string(b), err)
}

func main() {
	simple()
	simple1()
	simple2()
}
