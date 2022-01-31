package controller

import (
	"fmt"
	"go-smaole/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Data    interface{}
	Status  string
	Message string
}

func GetTeacher(c *gin.Context) {
	var result = Result{}
	result.Status = "0000"
	result.Message = "OK"
	result.Data = services.GetTeacher()

	c.JSON(http.StatusOK, result)
}

type teacher struct {
	Name string
	Age  int
}

func CreateTeacher(c *gin.Context) {
	var result = Result{}
	result.Status = "0000"
	result.Message = "OK"

	json := teacher{}
	c.BindJSON(&json)

	fmt.Println(json)
	result.Data = services.CreateTeacher(json.Name, json.Age)
	c.JSON(http.StatusOK, result)
}

func UpdateTeacher(c *gin.Context) {
	var result = Result{}
	result.Status = "0000"
	result.Message = "OK"
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		fmt.Println(id, err)
	}
	json := teacher{}
	c.BindJSON(&json)

	fmt.Println(json)
	services.UpdateTeacher(json.Name, json.Age, id)
	c.JSON(http.StatusOK, result)
}

func DeleteTeacher(c *gin.Context) {
	var result = Result{}
	result.Status = "0000"
	result.Message = "OK"
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		fmt.Println(id, err)
	}
	services.DeleteTeacher(id)
	c.JSON(http.StatusOK, result)
}
