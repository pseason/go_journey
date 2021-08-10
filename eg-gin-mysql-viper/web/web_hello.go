package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"springmars.com/gin-mysql-viper/model"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, model.SuccessResponseData("hello gin-mysql-viper"))
}
