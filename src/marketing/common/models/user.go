package models

import (
	"marketing/common/database"
)

type User struct {
	UserID int64
	Username string
	Password string
	Rid uint
}

// 获取用户ID
func GetUserRole(username string) (user User) {
	db, _ := database.Connect()
	defer db.Close()
	db.Table("t_user").Where("username = ?", username).First(&user)
	return 
}

// 认证检查
func CheckAuth(username, password string) bool {
	db, _ := database.Connect()
	defer db.Close()
	var user User
	db.Table("t_user").Select("username").Where("username = ? and password = ?", username, password).First(&user)
	if user.Username != "" {
		return true
	}
	return false
}
