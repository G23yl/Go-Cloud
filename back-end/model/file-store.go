// file-store结构
package model

import "gorm.io/gorm"

type FileStore struct {
	gorm.Model
	UserID      uint  `json:"user_id"`
	CurrentSize int64 `json:"current_size"`
	MaxSize     int64 `json:"max_size"`
}

func (f *FileStore) TableName() string {
	return "file_store"
}
