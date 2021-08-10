/*
线程安全单列模式
*/

package method

import "sync"

type BSingleton struct {
}

var sb *BSingleton

var lock sync.Mutex

// 利用 Sync.Mutex 进行加锁保证线程安全，但由于每次调用该方法都进行了加锁操作，在性能上不是很高效
func GetBSingletonLock() *BSingleton {
	lock.Lock()
	defer lock.Unlock()
	if sb == nil {
		sb = &BSingleton{}
	}
	return sb
}

//双重检查: 在懒汉式（线程安全）的基础上再进行优化，减少加锁的操作，保证线程安全的同时不影响性能
func GetBSingleton2Lock() *BSingleton {
	if sb == nil {
		lock.Lock()
		if sb == nil {
			sb = &BSingleton{}
		}
		lock.Unlock()
	}
	return sb
}
