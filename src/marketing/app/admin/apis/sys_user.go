package apis

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"marketing/app/admin/service"
	"marketing/common/utils/jwt"
)

// Login
// @Summary 登录
// @Description 用户认证
// @Tags 个人中心
// @Success 20000 {object} response.Response "{"code": 20000, "token": string}"
// @Router /api/user/login [post]
// @Security 
func Login(c *gin.Context) {
	json := make(map[string]interface{}) // 接收结构的内容
	c.BindJSON(&json)
	token := service.GetAuthToken(json["username"].(string), json["password"].(string))
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"message": "登录成功",
			"token": token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 30000,
			"message": "登录失败",
		})
	}
}

// GetInfo
// @Summary 获取信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 20000 {object} response.Response "{"code": 20000, "data": {...}}"
// @Router /api/user/getinfo [get]
// @Security 
func GetInfo(c *gin.Context) {
	// 解开请求头 token 角色id
	// 获取角色
	// 通过角色查找数据库角色名称
	headerMap := c.Request.Header
	token, _ := jwt.ParseToken(headerMap["X-Token"][0])
	if token != nil {
		data := make(map[string]interface{})
		data["roles"] = service.RoleNameFromId(token.RoleID)
		data["introduction"] = "内部用户"
		data["avatar"] = "#"
		data["name"] = token.Username
		c.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"message": "获取信息成功",
			"data": data,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 30000,
			"message": "X-Token信息异常",
		})
	}
}

// Logout
// @Summary 退出
// @Description 清除Token
// @Tags 个人中心
// @Success 20000 {object} response.Response "{"code": 20000, "message": string}"
// @Router /api/user/logout [post]
// @Security 
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"message": "退出成功",
	})
}
