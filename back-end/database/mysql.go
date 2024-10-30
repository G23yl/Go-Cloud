// 数据库操作封装
package database

import (
	"disk/config"
	"disk/model"
	"disk/model/response"
	"disk/utils/util_file"
	"fmt"
	"log"
	"strings"

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

func DBFindByID(userID uint) *model.User {
	var user model.User
	if db.Where("ID = ?", userID).First(&user).RowsAffected == 0 {
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
			MaxSize:     100 * 1024 * 1024, // 100M
		}
		if err := tx.Create(store).Error; err != nil {
			return err
		}
		return nil
	})
}

// 查询仓库容量和大小
func GetStoreInfo(userID uint) (storeID uint, capacity, size int64) {
	var store model.FileStore
	db.Where("user_id = ?", userID).First(&store)
	return store.ID, store.MaxSize, store.CurrentSize
}

// 查询各类文件的数量
func GetFilesNum(storeID uint) (*model.FilesNum, int64) {
	var files []model.File
	db.Where("file_store_id = ?", storeID).Find(&files)
	filesNum := &model.FilesNum{}
	for _, file := range files {
		switch file.Type {
		case 1:
			filesNum.DocNum++
		case 2:
			filesNum.ImgNum++
		case 3:
			filesNum.VideoNum++
		case 4:
			filesNum.MusicNum++
		case 5:
			filesNum.OtherNum++
		}
	}
	return filesNum, int64(len(files))
}

func GetFoldersNum(storeID uint) int64 {
	var folders []model.Folder
	db.Where("file_store_id = ?", storeID).Find(&folders)
	return int64(len(folders))
}

func GetCategoryFiles(storeID uint, fileType int) []model.File {
	var files []model.File
	db.Where("file_store_id = ? AND type = ?", storeID, fileType).Find(&files)
	return files
}

func UpdateAvatar(user *model.User) error {
	return db.Save(&user).Error
}

func CreateFile(storeID uint, filePath, fileName string, size int64) error {
	// 采用事务，防止出现文件保存了，但是仓库大小没更新的情况
	return db.Transaction(func(tx *gorm.DB) error {
		arr := strings.Split(fileName, ".")
		suffix := "." + arr[1]
		err := db.Create(&model.File{
			FileStoreID: storeID,
			FileName:    fileName,
			FilePath:    filePath,
			FileSize:    size,
			Type:        util_file.GetFileType(suffix),
		}).Error
		if err != nil {
			return err
		}
		// 更新仓库的大小
		var store model.FileStore
		db.Where("ID = ?", storeID).First(&store)
		store.CurrentSize += size
		return db.Save(&store).Error
	})
}

func GetInPathFiles(path string) ([]model.File, []model.Folder) {
	var files []model.File
	var folders []model.Folder
	db.Where("file_path = ?", path).Find(&files)
	db.Where("file_path = ?", path).Find(&folders)
	return files, folders
}

func CheckFileExists(storeID uint, filename, filepath string) bool {
	return db.Where("file_store_id = ? and file_name = ? and file_path = ?", storeID, filename, filepath).RowsAffected != 0
}
