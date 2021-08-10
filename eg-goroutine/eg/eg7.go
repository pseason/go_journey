package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	exit := make(chan int)

	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		exit <- 0
	})

	<-exit
	fmt.Println("main processor over")

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)

}
