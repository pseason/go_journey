package main

import (
	"fmt"

	"springmars.com/lock/m"
)

func main() {
	m.Add4Unlock()
	fmt.Println("-------------")
	m.Add4Lock()
	fmt.Println("-------------")
	m.Lock4Goroutine()
	fmt.Println("-------------")
	m.ReadWirte()
}
