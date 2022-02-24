package router

import (
	"github.com/gin-gonic/gin"
	"ptc/api/v1/account"
	"ptc/api/v1/mainpage"
	"ptc/api/v1/personalpage"
	"ptc/api/v1/upload"
	"ptc/internal/middleware"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(.AppMode)

	r := gin.Default()
	// 发送模块
	routerV1MainPage := r.Group("/v1/mainPage")
	{
		routerV1MainPage.POST("/send", middleware.JwtMiddleware(), mainpage.Send)
		routerV1MainPage.POST("/forward", middleware.JwtMiddleware(), mainpage.Forward)
		routerV1MainPage.GET("/page", mainpage.ShowPage)
		routerV1MainPage.POST("/like", mainpage.AddLike)
		routerV1MainPage.POST("/cancelLike", mainpage.CancelLike)
		routerV1MainPage.GET("/comment", mainpage.ShowComment)
		routerV1MainPage.GET("/moreComment", mainpage.ShowMoreComment)
		routerV1MainPage.POST("/insertComment", mainpage.InsertComment)
		routerV1MainPage.GET("/showCollection", mainpage.ShowCollection)
		routerV1MainPage.POST("/addCollection", mainpage.AddCollection)
	}
	// 登陆注册模块
	routerV1Account := r.Group("/v1/account")
	{
		routerV1Account.POST("/login", account.Login)
		routerV1Account.POST("/register", account.Register)
	}

	// 上传模块
	routerV1UpLoad := r.Group("/v1/upload")
	{
		routerV1UpLoad.POST("/post/image", upload.PostImage)
		routerV1UpLoad.POST("/profile/image", upload.ProfileImage)
	}

	// 个人主页模块
	routerV1PersonalPage := r.Group("/v1/personalPage")
	{
		routerV1PersonalPage.GET("/showPersonInfo", personalpage.ShowPersonalInfo)
		routerV1PersonalPage.GET("/showPersonalPost", personalpage.ShowPersonalPost)
		routerV1PersonalPage.GET("/showCollection", mainpage.ShowCollection)
		routerV1PersonalPage.GET("/showCollectionPost", personalpage.ShowCollectionPost)
		routerV1PersonalPage.POST("/newCollection", personalpage.NewCollection)
		routerV1PersonalPage.GET("/showProfile", personalpage.ShowProfile)
		routerV1PersonalPage.POST("/modifyProfile", personalpage.ModifyProfile)
		routerV1PersonalPage.POST("/modifyPhoto", personalpage.ModifyPhoto)
		routerV1PersonalPage.POST("/deletePhoto", personalpage.DeletePhoto)
		routerV1PersonalPage.POST("/modifyPassword", personalpage.ModifyPassword)
	}

	//注册zap日志框架的中间件
	return r
}
