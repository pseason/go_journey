package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"springmars.com/gin-air/model"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, model.SuccessResponseData("hello gin-gorm-mysql-viper ........."))
}
