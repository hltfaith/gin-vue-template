package service

import (
	"fmt"
	"crypto/md5"

	"marketing/common/models"
	"marketing/common/utils/jwt"
)

/**
* @FuncName 获取认证Token
* @Describe 
* @Param username 用户名 password 密码
*
* @return string格式
*/
func GetAuthToken(username, password string) string {
	isExist := models.CheckAuth(username, password)
	if isExist {
		user := models.GetUserRole(username) // 获取用户id
		token, err := jwt.GenerateToken(&user)
		if err == nil {
			return token
		}
		return ""
	}
	return ""
}

// 密码加盐
func HashSalt(password string) string {
	// md5加密
	var salt string = "marketing"
	var str string = password + salt
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

/**
* @FuncName 获取认证Token
* @Describe 
* @Param username 用户名 password 密码
*
* @return string格式
*/
