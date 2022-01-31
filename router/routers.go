package router

import (
	"go-smaole/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/teacher", controller.GetTeacher)
	r.POST("/api/teacher", controller.CreateTeacher)
	r.PUT("/api/teacher/:id", controller.UpdateTeacher)
	r.DELETE("/api/teacher/:id", controller.DeleteTeacher)

	return r
}
