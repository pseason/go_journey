package model

type Student struct {
	Id   int
	Name string
	Age  int
}

func NewStudent() *Student {
	return new(Student)
}
