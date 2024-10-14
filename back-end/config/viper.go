package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() {
	vp := viper.New()

	//设置配置文件名
	vp.SetConfigFile("config.yaml")
	// 读取配置文件
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal("viper读取配置文件失败...", err)
	}

	// 配置实时监控配置文件改动
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变...", in.Name)
		// 更新到全局Config里面
		if err := vp.Unmarshal(&config); err != nil {
			log.Fatal("viper更新配置失败...", err)
		}
	})
	if err := vp.Unmarshal(&config); err != nil {
		log.Fatal("viper配置失败...", err)
	}
	fmt.Printf("%#v\n", config)
}
