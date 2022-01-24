package main

import (
	"github.com/spf13/viper"
	"ptc/config"
	"ptc/internal/respository"
	"ptc/internal/router"
)

func main() {
	// 初始化配置文件
	config.InitConfig()
	// 初始化log
	//log.InitLogger()
	// 数据库初始化
	respository.InitDbContext()
	// 初始化gin的路由
	ginRoute := router.InitRouter()
	//运行
	port := viper.GetString("server.HttpPort")
	err := ginRoute.Run(":" + port)
	if err != nil {
		return
	}
}
