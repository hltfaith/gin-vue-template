package apis

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	
	"marketing/app/admin/service"
	"marketing/common/utils/jwt"
)

// GetRouter
// @Summary 获取路由表信息
// @Description 提供前端路由表
// @Tags 用户认证
// @Success 20000 {object} response.Response "{"code": 20000, "data": [...]}"
// @Router /api/role/router [get]
// @Security 
func GetRouter(c *gin.Context) {
	headerMap := c.Request.Header
	token, _ := jwt.ParseToken(headerMap["X-Token"][0])
	data := service.RouterInfo(token.RoleID)
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"message": "获取路由表成功",
		"data": data,
	})
}
