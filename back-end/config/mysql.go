package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSql() {
	dsn := GetMysqlDsn()

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接mysql失败...", err)
	}
	DB = _db
	fmt.Println("连接mysql成功!!!")
	// db.AutoMigrate(&model.User{}, &model.FileStore{}, &model.FileFolder{}, &model.File{})
}
