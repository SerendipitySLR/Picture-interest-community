package router

import (
	"Picture-interest-community/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/login",controller.Login)
	r.POST("/register",controller.Register)

	return r
}
