/*
接口嵌套
*/
package main

import "fmt"

type Walker interface {
	Walk()
}

type Flyer interface {
	Fly()
}

type Bird interface {
	Walker
	Flyer
}

type BigBird struct {
}

func (bb *BigBird) Walk() {
	fmt.Println("BigBird Walk")
}

func (bb *BigBird) Fly() {
	fmt.Println("BigBird Flyer")
}

func main() {
	var b Bird = new(BigBird)

	b.Fly()
	b.Walk()

}
