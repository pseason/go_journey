/*
singleton factory
*/

package method

type ASingleton struct {
}

var sa *ASingleton

func GetASingleton() *ASingleton {
	if sa == nil {
		sa = &ASingleton{}
	}
	return sa
}
