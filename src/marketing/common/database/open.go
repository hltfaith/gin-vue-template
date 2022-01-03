package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/manage_system_db?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)	// 全局禁用表名复数
	return db, err
}
