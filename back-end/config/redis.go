package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	// 配置redis
	options := GetRedisOptions()
	Client := redis.NewClient(options)

	// 测试redis是否连接
	if _, err := Client.Ping(ctx).Result(); err != nil {
		log.Fatal("连接redis失败...", err)
	}
	fmt.Println("连接redis成功!!!")
}
