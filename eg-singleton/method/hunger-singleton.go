/*
饿汉-单列模式
*/
package method

type HSingleton struct {
}

var sh *HSingleton = &HSingleton{}

func GetHSingleton() *HSingleton {
	return sh
}
