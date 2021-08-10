package main

import (
	"springmars.com/pingo-provider/rpc"

	"github.com/dullgiulio/pingo"
)

func main() {

	pingo.Register(rpc.NewMyPlugin())

	pingo.Run()
}
