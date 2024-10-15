package main

import (
	"disk/config"
	"disk/database"
	"disk/router"
	"log"
)

func main() {
	// 初始化viper，加载配置
	config.InitViper()
	// 初始化数据库，mysql和redis
	database.ConnectSql()
	database.ConnectRedis()
	// 获取路由
	router := router.SetUpRouter()

	if err := router.Run(config.GetServerPort()); err != nil {
		log.Fatal("服务器启动失败...", err)
	}
}
