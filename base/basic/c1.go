package main

import (
	"fmt"
	"math"
)

/**

hello go

*/
func main() {
	var s string = "hello world"
	var boo bool = true
	var i int = 8
	var f float64 = .8454
	var ff float32 = 16777216
	fmt.Println(s, boo, i, f, ff == 1)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.3f\n", math.Pi)
	fmt.Printf("%.7f\n", f)
	//复数
	var fx complex128 = complex(1, 2)
	var fy complex128 = complex(3, 4)
	fmt.Println(fx)
	fmt.Println(fy)
	fmt.Println(fx * fy)
	fmt.Println(real(fx * fy))
	fmt.Println(imag(fx * fy))
}
