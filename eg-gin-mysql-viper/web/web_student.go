package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"springmars.com/gin-mysql-viper/model"
	"springmars.com/gin-mysql-viper/service"
)

// find/:id
func FindStudentById(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	fmt.Println("FindStudentById:", id)

	res := service.FindStudentById(id)
	c.JSON(http.StatusOK, res)
}

// 绑定QueryString示例 (/findAll?pageSize=1&pageNum=10)
func FindAllStudent(c *gin.Context) {
	var pageInfo model.PageInfo
	var res model.ResponseData
	if err := c.ShouldBind(&pageInfo); err == nil {
		fmt.Println("FindAllStudent", pageInfo)
		res = service.FindAllStudent(pageInfo.PageSize, pageInfo.PageNum)
	} else {
		res = model.ErrorResponseData(err)
	}
	c.JSON(http.StatusOK, res)
}

// 绑定JSON的示例 ({"name": "lucy", "age": "16"})
func InsertStudent(c *gin.Context) {
	var stu model.Student
	var res model.ResponseData
	if err := c.BindJSON(&stu); err == nil {
		fmt.Println("FindAllStudent", stu)
		res = service.InsertStudent(stu)
	} else {
		res = model.ErrorResponseData(err)
	}
	c.JSON(http.StatusOK, res)
}

// 绑定JSON的示例 ({"id": 1 ,"name": "lucy", "age": "16"})
func UpdateStudent(c *gin.Context) {
	fmt.Println("------------------")

	var stu model.Student
	var res model.ResponseData
	if err := c.BindJSON(&stu); err == nil {
		fmt.Println("UpdateStudent", stu)
		res = service.UpdateStudent(stu)
	} else {
		res = model.ErrorResponseData(err)
	}
	c.JSON(http.StatusOK, res)
}

func DeleteStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	fmt.Println("DeleteStudent:", id)

	res := service.DeleteStudent(id)
	c.JSON(http.StatusOK, res)
}
