/*
错误接口
*/
package main

import (
	"fmt"
)

type GameError struct {
	Code  int
	Cause string
}

func (ge GameError) Error() string {
	return fmt.Sprintf("Error Code: %d, cause:\n	%s ", ge.Code, ge.Cause)
}

func simple(k interface{}) error {
	switch k.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	default:
		return GameError{Code: 501, Cause: "类型未匹配"}
	}
	return nil
}

func main() {
	err := simple(1)
	fmt.Println(err)
	err = simple(new(GameError))
	fmt.Println(err)

	var inter1 interface{} = 1

	var inter interface{}

	fmt.Println(inter, inter1)
}
