/*
雪花算法生成唯一keyId
*/
package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("snowflake.NewNode error: ", err)
		goto ELable
	}
	for i := 0; i < 3; i++ {
		id := node.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
			"\n",
		)
	}

ELable:
	os.Exit(1)

}
