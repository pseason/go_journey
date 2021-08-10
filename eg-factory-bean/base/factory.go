/*

 */
package base

// 类接
type Class interface {
	// do something
	Do()
}

var (
	// 保存注册好的工厂信息
	factoryByName = make(map[string]func() Class, 6)
)

// 注册
func Register(name string, callback func() Class) {
	factoryByName[name] = callback
}

func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
