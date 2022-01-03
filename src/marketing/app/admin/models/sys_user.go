package models

import (
	"github.com/jinzhu/gorm"
)

// 用户表
type User struct {
	gorm.Model
	UserID string `gorm:"not null;unique_index;comment:'用户ID'"`
	Username string	`gorm:"not null;comment:'账号'"`
	Password string `gorm:"not null;comment:'密码'"`
	Rid uint `gorm:"not null;comment:'角色ID'"`
	Role Role `gorm:"foreignKey:Rid;references:RoleID"`
}

func (User) TableName() string {
	return "t_user"
}

// 用户信息表
type UserProfile struct {
	gorm.Model
	Uid string `gorm:"not null;unique_index;comment:'用户ID'"`
	Nickname string `gorm:"comment:'昵称'"`
	Email string `gorm:"comment:'邮箱'"`
	User User `gorm:"references:UserId"`
}

func (UserProfile) TableName() string {
	return "t_user_profile"
}
