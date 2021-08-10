package service

import (
	"springmars.com/gin-gorm-mysql-viper/entity"
	"springmars.com/gin-gorm-mysql-viper/model"
	"springmars.com/gin-gorm-mysql-viper/repo"
)

// find by id
func FindStudentById(id int) model.ResponseData {
	student, err := repo.FindStudentById(id)
	if err != nil {
		return model.ErrorResponseData(err)
	}
	return model.SuccessResponseData(*student)
}

// find all
func FindAllStudent(pageSize, pageNum int) model.ResponseData {
	students, err := repo.FindAllStudent(pageSize, pageNum)
	if err != nil {
		return model.ErrorResponseData(err)
	}
	return model.SuccessResponseData(students)
}

// insert
func InsertStudent(stu entity.Student) model.ResponseData {
	id, err := repo.InsertStudent(stu)
	if err != nil {
		return model.ErrorResponseData(err)
	}
	return model.SuccessResponseData(id)
}

// update
func UpdateStudent(stu entity.Student) model.ResponseData {
	cnt, err := repo.UpdateStudent(stu)
	if err != nil {
		return model.ErrorResponseData(err)
	}
	return model.SuccessResponseData(cnt)
}

// delete
func DeleteStudent(id int) model.ResponseData {
	cnt, err := repo.DeleteStudent(id)
	if err != nil {
		return model.ErrorResponseData(err)
	}
	return model.SuccessResponseData(cnt)
}
