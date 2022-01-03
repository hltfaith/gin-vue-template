package service

import (
	// "fmt"

	"marketing/common/models"
)

/**
* @FuncName 获取路由信息
* @Describe 按照 vue router js格式返回
* @Param role 角色标识
* 
* @return 数组格式
*/
func RouterInfo(rid uint) []map[string]interface{} {
	router := models.GetPermissionFormRid(rid)
	routerList := make([]map[string]interface{},0)
	menuMap := models.GetMenu()
	// 单页面
	pageList := make([]string,0)
	pageList = menuMap["page"].([]string)
	for _, menu := range pageList{
		for _, r := range router {
			if r.Redirect == menu {
				singleMap := make(map[string]interface{})
				singleMap["path"] = r.Path
				singleMap["component"] = r.Component
				singleMap["redirect"] = r.Redirect
				
				childrenList := make([]map[string]interface{},0)
				childrenMap := make(map[string]interface{})
				meta := make(map[string]string)
				childrenMap["path"] = r.SubPath
				childrenMap["name"] = r.Title
				childrenMap["component"] = r.SubComponent
				meta["title"] = r.Title
				meta["icon"] = r.Icon
				childrenMap["meta"] = meta
				childrenList = append(childrenList, childrenMap)
				singleMap["children"] = childrenList

				if len(singleMap) != 0 {
					routerList = append(routerList, singleMap)
				}
			}
		}
	}
	// 组页面
	groupList := make([]string,0)
	groupList = menuMap["group"].([]string)
	for _, group := range groupList {
		childrenList := make([]map[string]interface{},0)
		groupMap := make(map[string]interface{})
		for _, r := range router {
			if r.IsGroup == "是" && r.Path == group {
				subChildren := make(map[string]interface{})
				meta := make(map[string]interface{})
				subMeta := make(map[string]interface{})

				groupMap["path"] = r.Path
				groupMap["component"] = r.Component
				groupMap["redirect"] = r.Redirect
				groupMap["name"] = r.Title
				meta["title"] = r.Title
				meta["icon"] = r.Icon
				groupMap["meta"] = meta
				
				subChildren["path"] = r.SubPath
				subChildren["name"] = r.SubTitle
				subChildren["component"] = r.SubComponent
				subMeta["title"] = r.SubTitle
				subMeta["icon"] = r.SubIcon
				subChildren["meta"] = subMeta
				childrenList = append(childrenList, subChildren)
				groupMap["children"] = childrenList
			}
		}
		if len(groupMap) != 0 {
			routerList = append(routerList, groupMap)
		}
	}
	return routerList
}

/**
* @FuncName 获取角色名称
* 
* @return 返回角色标识, 角色名称
*/
func RoleNameFromId(rid uint) []string {
	// 角色标识 
	name := models.GetRoleIDFromRId(&rid)
	// 角色名称
	return []string{name}
}
