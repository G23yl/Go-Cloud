// 数据库操作封装
package database

import (
	"disk/config"
	"disk/model"
	"disk/model/response"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectSql() {
	dsn := config.GetMysqlDsn()

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接mysql失败...", err)
	}
	db = _db
	fmt.Println("连接mysql成功!!!")
	if err := db.AutoMigrate(&model.User{}, &model.FileStore{}, &model.Folder{}, &model.File{}); err != nil {
		log.Fatal("数据库迁移失败...", err)
	}
	fmt.Println("数据库迁移成功!!!")
}

func DBFindByUsername(username string) *model.User {
	var user model.User
	if db.Where("username = ?", username).First(&user).RowsAffected == 0 {
		return nil
	}
	return &user
}

// 检验用户名和邮箱是否存在
func CheckUsernameAndEmail(username, email string) *response.Error {
	var user model.User
	if db.Where("username = ?", username).First(&user).RowsAffected != 0 {
		return &response.ErrorUsernameExist
	}
	if db.Where("email = ?", email).First(&user).RowsAffected != 0 {
		return &response.ErrorEmailExist
	}
	return nil
}

// 创建用户和文件仓库
func DBCreateUserAndFileStore(user *model.User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		store := &model.FileStore{
			UserID:      user.ID,
			CurrentSize: 0,
			MaxSize:     500 * 1024 * 1024, // 500M
		}
		if err := tx.Create(store).Error; err != nil {
			return err
		}
		return nil
	})
}
