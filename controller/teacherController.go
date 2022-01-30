package controller

import (
	"go-smaole/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Data    interface{}
	Status  string
	Message string
}

func GetUserName(c *gin.Context) {
	var result = Result{}
	result.Status = "0000"
	result.Message = "OK"
	result.Data = services.Get()

	c.JSON(http.StatusOK, result)
}
