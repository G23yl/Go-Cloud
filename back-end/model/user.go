// user结构
package model

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	Avatar   string `json:"avatar"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

// 设置表名
func (u *User) TableName() string {
	return "user"
}
