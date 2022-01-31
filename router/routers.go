package router

import (
	"go-smaole/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/teacher", controller.GetTeacher)
	r.POST("/api/teacher", controller.CreateTeacher)
	return r
}
