package model

type Student struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}

func NewStudent() *Student {
	return new(Student)
}
