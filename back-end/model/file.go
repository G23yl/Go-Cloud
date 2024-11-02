// file结构
package model

import "gorm.io/gorm"

// 文件结构定义
type File struct {
	gorm.Model
	FileStoreID    uint   `json:"file_store_id"`    // 所属文件仓库id
	ParentFolderID uint   `json:"parent_folder_id"` // 父文件夹id
	FileName       string `json:"file_name"`        // 文件名
	FilePath       string `json:"file_path"`        // 文件存储的路径
	DownloadCount  int64  `json:"download_count"`   // 下载次数
	FileSize       int64  `json:"file_size"`        // 文件大小
	Type           int    `json:"type"`             // 文件类型  1,2,3,4,5
}

func (f *File) TableName() string {
	return "file"
}

// 各种类型文件数量
type FilesNum struct {
	DocNum   int64
	ImgNum   int64
	VideoNum int64
	MusicNum int64
	OtherNum int64
}
