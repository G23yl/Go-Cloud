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
		// 获取父文件夹ID
		var parentFolderID uint = 0
		if filePath != "/" {
			parentFolderName, parentFolderPath := util_file.GetParentFolderNameAndPath(filePath)
			var folder model.Folder
			tx.Where("file_store_id = ? and file_path = ? and folder_name = ?", storeID, parentFolderPath, parentFolderName).First(&folder)
			parentFolderID = folder.ID
		}
		arr := strings.Split(fileName, ".")
		suffix := "." + arr[1]
		err := tx.Create(&model.File{
			FileStoreID:    storeID,
			FileName:       fileName,
			FilePath:       filePath,
			FileSize:       size,
			Type:           util_file.GetFileType(suffix),
			ParentFolderID: parentFolderID,
		}).Error
		if err != nil {
			return err
		}
		// 更新仓库的大小
		var store model.FileStore
		tx.Where("ID = ?", storeID).First(&store)
		store.CurrentSize += size
		return tx.Save(&store).Error
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
	var file model.File
	return db.Where("file_store_id = ? and file_name = ? and file_path = ?", storeID, filename, filepath).First(&file).RowsAffected != 0
}

func DeleteFile(storeID uint, fileID uint) error {
	// 要用事务
	return db.Transaction(func(tx *gorm.DB) error {
		var file model.File
		tx.Where("file_store_id = ? and ID = ?", storeID, fileID).First(&file)
		err := tx.Where("file_store_id = ? and ID = ?", storeID, fileID).Delete(&model.File{}).Error
		if err != nil {
			return err
		}
		// 更新仓库的大小
		var store model.FileStore
		tx.Where("ID = ?", storeID).First(&store)
		store.CurrentSize -= file.FileSize
		//FIXME 这个地方应该有bug，但是不知道为什么，以后再修复吧
		if store.CurrentSize < 0 {
			store.CurrentSize = 0
		}
		return tx.Save(&store).Error
	})
}

func DeleteFolders(storeID, folderID uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// BFS递归删除
		deletedSt := make([]model.Folder, 0)
		var folder model.Folder
		tx.Where("file_store_id = ? and ID = ?", storeID, folderID).First(&folder)
		deletedSt = append(deletedSt, folder)
		for len(deletedSt) != 0 {
			f := deletedSt[0]
			deletedSt = deletedSt[1:]
			// 删掉f文件夹
			tx.Where("file_store_id = ?", storeID).Delete(&f)
			// 找出parent_folder_id = f.ID的文件
			var files []model.File
			tx.Where("file_store_id = ? and parent_folder_id = ?", storeID, f.ID).Find(&files)
			for _, file := range files {
				// 调用删除文件的函数，同时删除文件和更新仓库大小
				DeleteFile(storeID, file.ID)
			}
			// 找出parent_folder_id = f.ID的文件夹
			var folders []model.Folder
			tx.Where("file_store_id = ? and parent_folder_id = ?", storeID, f.ID).Find(&folders)
			deletedSt = append(deletedSt, folders...)
		}
		return nil
	})
}

// 创建文件夹
func CreateFolder(storeID uint, path, folderName string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var parentFolderID uint = 0
		if path != "/" {
			parentFolderName, parentFolderPath := util_file.GetParentFolderNameAndPath(path)
			var folder model.Folder
			tx.Where("file_store_id = ? and file_path = ? and folder_name = ?", storeID, parentFolderPath, parentFolderName).First(&folder)
			parentFolderID = folder.ID
		}
		return tx.Create(&model.Folder{
			FileStoreID:    storeID,
			FilePath:       path,
			FolderName:     folderName,
			ParentFolderID: parentFolderID,
		}).Error
	})
}

// 下载文件更新文件下载次数
// func DownloadCountUpdate(storeID, fileID uint) {
// 	var file model.File
// 	db.Where("file_store_id = ? and ID = ?", storeID, fileID).First(&file)
// 	file.DownloadCount++
// 	db.Save(&file)
// }
