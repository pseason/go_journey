/*
once singelton
*/
package method

import "sync"

type OSingleton struct {
}

var so *OSingleton

var once sync.Once

func GetOSingletonOnce() *OSingleton {
	once.Do(func() {
		so = new(OSingleton)
	})
	return so
}
