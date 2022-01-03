package models

import (
	// "fmt"

	"marketing/app/admin/models"
	"marketing/common/database"
)

/**
* @FuncName 获取角色标识
* @Describe 通过角色id查询角色名称
*
* @return 返回string格式
* 	admin
*/
func GetRoleIDFromRId(rid *uint) string {
	db, _ := database.Connect()
	defer db.Close()
	var dbModel models.Role
	db.Table("t_role").Select("role_id").First(&dbModel, "rid = ?", rid)
	return dbModel.RoleID
}

/**
* @FuncName 获取所有角色路由
* 
* @return 返回数组格式
* 	["admin", "user"]
*/
func GetRolePermission(pid uint) []string {
	db, _ := database.Connect()
	defer db.Close()
	// 取角色编号 [1, 2]
	var dbModel []models.RoleBePermission
	db.Table("t_role_permission").Select("rid").Where("pid = ?", pid).Find(&dbModel)
	var ridList []uint
	for _, item := range dbModel {
		ridList = append(ridList, item.Rid)
	}
	// 取角色标识 ["admin", "user"]
	var dbRole []models.Role
	db.Table("t_role").Select("role_id").Where("rid IN (?)", ridList).Find(&dbRole)
	var roleList []string
	for _, item := range dbRole {
		roleList = append(roleList, item.RoleID)
	}
	return roleList
}

/**
* @FuncName 获取所有路由权限
* @Describe 通过角色id查询角色名称
* @Param  
*
* @return 返回models数组
*/
func GetPermissionFormRid(rid uint) (permissions []models.Permission) {
	db, _ := database.Connect()
	defer db.Close()
	// 取pid编号 [1, 2]
	var dbModel []models.RoleBePermission
	db.Table("t_role_permission").Select("pid").Where("rid = ?", rid).Find(&dbModel)
	var pidList []uint
	for _, item := range dbModel {
		pidList = append(pidList, item.Pid)
	}
	db.Table("t_permission").Where("permission_id IN (?)", pidList).Find(&permissions)
	return 
}

/**
* @FuncName 获取所有路由
* 
*/
func GetPermission() (permissions []models.Permission) {
	db, _ := database.Connect()
	defer db.Close()
	db.Find(&permissions)
	return 
}

/**
* @FuncName 获取菜单列表
* 
* @return 返回map格式
* 	{"page": ["/dashboard"], "group": ["/get/user", "/get/role"]}
*/
func GetMenu() map[string]interface{} {
	db, _ := database.Connect()
	defer db.Close()
	var groupPermission []models.Permission
	var pagePermission []models.Permission
	db.Table("t_permission").Select("redirect").Where("is_group = ?", "是").Find(&groupPermission)
	db.Table("t_permission").Select("redirect").Where("is_group = ?", "否").Find(&pagePermission)
	menuMap := make(map[string]interface{}, 0)
	groupList := make([]string, 0)
	pageList := make([]string, 0)
	for _, group := range groupPermission {
		groupList = append(groupList, group.Redirect)
	}
	for _, page := range pagePermission {
		pageList = append(pageList, page.Redirect)
	}
	menuMap["group"] = removeDuplicate(groupList)
	menuMap["page"] = removeDuplicate(pageList)
	return menuMap
}

/**
* @FuncName []string数组去重
* @Describe 通过map主键唯一的特性过滤重复元素
*
* @return []string数组格式
*/
func removeDuplicate(arr []string) []string {
    resArr := make([]string, 0)
    tmpMap := make(map[string]interface{})
    for _, val := range arr {
        //判断主键为val的map是否存在
        if _, ok := tmpMap[val]; !ok {
            resArr = append(resArr, val)
            tmpMap[val] = nil
        }
    } 
    return resArr
}
