/*
秒 转 分 时 天
*/

package main

import "fmt"

const (
	SecondPerMinute = 60
	SecondPerHour   = 60 * 60
	SecondPerDay    = 60 * 60 * 24
)

func translate(second int) (minute, hour, day int) {
	minute = second / SecondPerMinute
	hour = second / SecondPerHour
	day = second / SecondPerDay
	return
}

func main() {
	fmt.Println(translate(16660000))
	fmt.Println(translate(60000))
}
