// Redis API封装
package database

import (
	"context"
	"disk/config"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	// 配置redis
	options := config.GetRedisOptions()
	client = redis.NewClient(options)

	// 测试redis是否连接
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatal("连接redis失败...", err)
	}
	fmt.Println("连接redis成功!!!")
}

func Set(key, value string, duration time.Duration) error {
	return client.Set(ctx, key, value, duration).Err()
}

func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func Delete(key string) error {
	return client.Del(ctx, key).Err()
}
