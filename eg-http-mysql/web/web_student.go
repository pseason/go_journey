package web

import (
	"fmt"
	"net/http"
	"strconv"

	"springmars.com/http-mysql/model"
	"springmars.com/http-mysql/service"
)

func FindStudentById(w http.ResponseWriter, r *http.Request) {
	idstr := GetUrlArg(r, "id")
	id, _ := strconv.Atoi(idstr)

	fmt.Println("FindStudentById:", id)

	res := service.FindStudentById(id)
	WriteJson(w, res)
}

func FindAllStudent(w http.ResponseWriter, r *http.Request) {
	ps := GetUrlArg(r, "pageSize")
	pn := GetUrlArg(r, "pageNum")
	pageSize, _ := strconv.Atoi(ps)
	pageNum, _ := strconv.Atoi(pn)

	fmt.Println("FindAllStudent")

	res := service.FindAllStudent(pageSize, pageNum)
	WriteJson(w, res)
}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var stu model.Student
	GetBodyArg(r, &stu)

	fmt.Println("InsertStudent", stu)

	res := service.InsertStudent(stu)
	WriteJson(w, res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var stu model.Student
	GetBodyArg(r, &stu)

	fmt.Println("UpdateStudent", stu)

	res := service.UpdateStudent(stu)
	WriteJson(w, res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idstr := GetUrlArg(r, "id")
	id, _ := strconv.Atoi(idstr)

	fmt.Println("DeleteStudent:", id)

	res := service.DeleteStudent(id)
	WriteJson(w, res)
}
