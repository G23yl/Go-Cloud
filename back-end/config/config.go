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
	Email struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Address  string `mapstructure:"addr"`
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

func GetServerHost() string {
	return config.Server.Host
}

func GetServerPort() string {
	return fmt.Sprintf(":%d", config.Server.Port)
}

func GetSecret() string {
	return config.JWT.Secret
}

func GetExpired() int {
	return config.JWT.Expired
}

func GetEmailFrom() string {
	return config.Email.Username
}

func GetEmailHost() string {
	return config.Email.Host
}

func GetEmailPort() int {
	return config.Email.Port
}

func GetEmailPassword() string {
	return config.Email.Password
}
