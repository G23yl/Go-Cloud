package router

import (
	"disk/middleware"
	"disk/services"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	// 设置跨域
	router.Use(middleware.Cors())
	// 设置静态资源地址映射
	// 头像图片映射
	router.Static("/images", "static/images")

	router.POST("/login", services.Login)
	router.POST("/register", services.SignUp)
	router.POST("/verify", services.Verify)

	user := router.Group("user")
	user.Use(middleware.Auth())
	user.Use(middleware.Cors())
	{
		user.POST("/avatar", services.ChangeAvatar)
	}

	cloud := router.Group("cloud")
	cloud.Use(middleware.Auth())
	// 设置跨域
	cloud.Use(middleware.Cors())
	{
		cloud.GET("/index", services.Index)
		cloud.GET("/docs", services.Docs)
		cloud.GET("/images", services.Images)
		cloud.GET("/videos", services.Videos)
		cloud.GET("/audios", services.Audios)
		cloud.GET("/others", services.Others)
		cloud.POST("/files", services.Upload)
	}
	return router
}
