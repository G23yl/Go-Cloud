// folder结构
package model

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	FileStoreID    uint   `json:"file_store_id"`
	ParentFolderID uint   `json:"parent_folder_id"`
	FolderName     string `json:"folder_name"`
}

func (f *Folder) TableName() string {
	return "folder"
}