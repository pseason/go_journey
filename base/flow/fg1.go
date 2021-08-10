/*
模拟聊天
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func chat() {
	// 命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name: ")
	str, err := inputReader.ReadString('\n')
	if err == nil {
		name := str[:len(str)-2]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	} else {
		fmt.Println("some error occur...", err)
		goto ELable
	}
	// 循环输入，达到条件退出
	for {
		str, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		str = str[:len(str)-2]
		// 全部转换为小写。
		str = strings.ToLower(str)
		switch str {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出。
			goto ELable
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
ELable:
	fmt.Println("programe exit...")
	os.Exit(1)
}

func main() {
	chat()
}
