package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Mysql struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Dbname      string `yaml:"dbname"`
		OtherConfig string `yaml:"other_config"`
	}
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Index    int    `yaml:"index"`
		Protocol int    `yaml:"protocol"`
	}
	JWT struct {
		Secret  string `yaml:"secret"`
		Expired int    `yaml:"expired"`
	}
}

var config Config

func GetMysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Dbname, config.Mysql.OtherConfig)
}

func GetRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Index,
		Protocol: config.Redis.Protocol,
	}
}

func GetServerPort() string {
	return fmt.Sprintf(":%d", config.Server.Port)
}
