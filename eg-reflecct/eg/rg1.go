package eg

import (
	"fmt"
	"reflect"
)

// 简单反射方法
func simple() {
	var a int
	typeOfa := reflect.TypeOf(a)
	fmt.Println(typeOfa)
	fmt.Println(typeOfa.Name())
	fmt.Println(typeOfa.Kind())

}

// 指针类型反射
type Cat struct {
	Name string
	// 带有结构体tag的字段
	Type int `json:"type" id:"100"`
}

func ptr() {
	cat := &Cat{}
	typeOfCat := reflect.TypeOf(cat)
	fmt.Println(typeOfCat.Name())
	fmt.Println(typeOfCat.Kind())

	typeOfCat = typeOfCat.Elem()
	fmt.Println(typeOfCat.Name())
	fmt.Println(typeOfCat.Kind())
}

func field() {
	cat := Cat{Name: "xiaohua", Type: 1}
	typeOfCat := reflect.TypeOf(cat)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	if fieldType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(fieldType.Tag.Get("json"), fieldType.Tag.Get("id"))
	}
}

func canSet() {
	cat := Cat{Name: "xiaohua", Type: 1}
	valueOfCat := reflect.ValueOf(cat)
	fmt.Println("field can set:", valueOfCat.CanSet())
	fmt.Println("field can set:", valueOfCat.Field(0).CanSet())
	fmt.Println("field can set:", valueOfCat.Field(1).CanSet())
	// 不能set
	// valueOfCat.Field(1).SetInt(5)
	// fmt.Println(cat)
	fmt.Println("--------------")
	cat1 := Cat{Name: "xiaohua", Type: 1}
	valueOfCat1 := reflect.ValueOf(&cat1)
	fmt.Println("field can set:", valueOfCat1.CanSet())
	valueOfCat1Elem := valueOfCat1.Elem()
	fmt.Println("field can set:", valueOfCat1Elem.CanSet())
	valueOfCat1Elem.Field(0).SetString("xiaogie")
	valueOfCat1Elem.Field(1).SetInt(3)
	fmt.Println(cat1)
}

func valid() {
	var a *int
	// 返回值是否为 nil
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	// 判断值是否有效
	fmt.Println("var a *int:", reflect.ValueOf(a).IsValid())
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	// *int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	// 实例化一个结构体
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

func new() {
	type dog struct {
		Name string
	}
	dg := dog{}
	dgType := reflect.TypeOf(dg)

	ins := reflect.New(dgType)
	fmt.Println(ins.Type(), ins.Kind())

	fmt.Println(ins.CanSet())

	dgg := ins.Interface().(*dog)
	dgg.Name = "xxxx"
	fmt.Println(*dgg)
}

func add(a, b int) int {
	return a + b
}

func callfunc() {
	valueOfadd := reflect.ValueOf(add)

	reflectList := []reflect.Value{reflect.ValueOf(12), reflect.ValueOf(122)}

	res := valueOfadd.Call(reflectList)
	fmt.Println("res:", res[0].Int())
}

func init() {
	simple()
	fmt.Println("---------------")
	ptr()
	fmt.Println("---------------")
	field()
	fmt.Println("---------------")
	canSet()
	fmt.Println("---------------")
	valid()
	fmt.Println("---------------")
	new()
	fmt.Println("---------------")
	callfunc()
}
