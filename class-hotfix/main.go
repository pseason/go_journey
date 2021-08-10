package main

import (
	"log"
	"springmars.com/class-hotfix/process"
)

func main() {
	list := process.JavaProcessList()
	for _, item := range list {
		log.Printf("process: pid:%d name:%s", item.PID, item.Name)
	}
}
