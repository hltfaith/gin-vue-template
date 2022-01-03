package models

import (
	"github.com/jinzhu/gorm"
)

// 角色表
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(128);not null;comment:'角色名称'"`
	RoleID string `gorm:"not null;unique_index;comment:'角色标识'"`
	Rid uint `gorm:"not null;comment:'角色ID'"`
	Remark string `gorm:"size:255;comment:'描述'"`
}

func (Role) TableName() string {
	return "t_role"
}

// 权限表
type Permission struct {
	gorm.Model
	PermissionId uint `gorm:"not null;unique_index;comment:'权限ID'"`
	IsGroup string `gorm:"not null;comment:'是否是组'"`
	
	Title string  `gorm:"not null;comment:'模块名称'"`
	Path string `gorm:"not null;comment:'模块路径'"`
	Redirect string `gorm:"not null;comment:'重定向路径'"`
	Component string `gorm:"not null;comment:'模块部件'"`
	Icon string `gorm:"not null;comment:'模块icon'"`
	
	SubTitle string `gorm:"comment:'子模块标题'"`
	SubComponent string `gorm:"comment:'子模块部件'"`
	SubPath string `gorm:"comment:'子模块路径'"`
	SubIcon string `gorm:"comment:'子模块icon'"`
}

func (Permission) TableName() string {
	return "t_permission"
}

// 角色权限关联表
type RoleBePermission struct {
	gorm.Model
	Pid uint `gorm:"not null;comment:'权限ID'"`
	Rid uint `gorm:"not null;comment:'角色ID'"`
}

func (RoleBePermission) TableName() string {
	return "t_role_permission"
}
