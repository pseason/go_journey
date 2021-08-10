/*
接口类型之间的转换
*/
package main

import "fmt"

type Walker interface {
	Walk()
}

type Flyer interface {
	Fly()
}

type Bird struct {
}

func (bb *Bird) Walk() {
	fmt.Println("Bird Walk")
}

func (bb *Bird) Fly() {
	fmt.Println("Bird Flyer")
}

type Pig struct {
}

func (p *Pig) Walk() {
	fmt.Println("Pig Walk")
}

func main() {
	mps := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}
	for name, obj := range mps {
		f, isFlyer := obj.(Flyer)
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}
}
