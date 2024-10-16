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

	cloud := router.Group("cloud")
	cloud.Use(middleware.Auth())

	// {
	// 	cloud.GET("/index", services.Index)
	// 	cloud.GET("/files", services.Files)
	// 	cloud.GET("/upload", services.Upload)
	// 	cloud.GET("/docs", services.Docs)
	// 	cloud.GET("/images", services.Images)
	// 	cloud.GET("/videos", services.Videos)
	// 	cloud.GET("/musics", services.Musics)
	// 	cloud.GET("/others", services.Others)
	// 	cloud.GET("/download", services.DownloadFile)
	// }
	// {
	// 	cloud.DELETE("/file", services.DeleteFile)
	// 	cloud.DELETE("/folder", services.DeleteFolder)
	// }
	// {
	// 	cloud.POST("/file", services.HandleUpload)
	// 	cloud.POST("/folder", services.AddFolder)
	// 	cloud.PUT("/folder", services.UpdateFolder)
	// }

	return router
}
