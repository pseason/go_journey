package repo

import (
	"springmars.com/gin-gorm-mysql-viper/entity"
)

// find by id
func FindStudentById(id int) (*entity.Student, error) {
	var stu entity.Student
	result := MysqlDb.First(&stu, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &stu, nil
}

// find all
func FindAllStudent(pageSize, pageNum int) (*[]entity.Student, error) {
	stus := make([]entity.Student, 0)
	result := MysqlDb.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&stus)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &stus, nil
}

// insert
func InsertStudent(stu entity.Student) (int, error) {
	result := MysqlDb.Create(&stu)
	if err := result.Error; err != nil {
		return 0, err
	}
	return stu.Id, nil
}

// update
func UpdateStudent(stu entity.Student) (int, error) {
	result := MysqlDb.Save(&stu)
	if err := result.Error; err != nil {
		return 0, err
	}
	return int(result.RowsAffected), nil
}

// delete
func DeleteStudent(id int) (int, error) {
	result := MysqlDb.Delete(&entity.Student{}, id)
	if err := result.Error; err != nil {
		return 0, err
	}
	return int(result.RowsAffected), nil
}
