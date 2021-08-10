package entity

type Student struct {
	Id   int    `form:"id" json:"id" gorm:"primaryKey"`
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}

// gorm 接口，如果不实现，TableName 将解析为students
func (Student) TableName() string {
	return "student"
}

func NewStudent() *Student {
	return new(Student)
}
