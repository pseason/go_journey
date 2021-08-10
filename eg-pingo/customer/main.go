package main

import (
	"fmt"

	"github.com/dullgiulio/pingo"
)

func main() {
	// 从创建的可执行文件中创建一个新插件。通过 TCP 连接到它
	p := pingo.NewPlugin("tcp", "../provider/pingo-provider")
	// 启动
	p.Start()
	// 使用完插件后停止它
	defer p.Stop()

	var resp string
	err := p.Call("MyPlugin.SayHello", "Goer", &resp)
	fmt.Println(resp, err)
}
