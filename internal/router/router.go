package router

import (
	"github.com/gin-gonic/gin"
	"ptc/api/v1/mainpage"
	"ptc/pkg/log"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(.AppMode)

	r := gin.Default()
	//发送模块
	routerV1MainPage := r.Group("/v1/mainPage")
	{
		routerV1MainPage.POST("/send", mainpage.Send)
		routerV1MainPage.POST("/follow", mainpage.Forward)
	}
	// 其他模块...

	//注册zap日志框架的中间件
	return r
}
