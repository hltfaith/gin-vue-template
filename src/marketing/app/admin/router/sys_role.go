package router

import (
	"github.com/gin-gonic/gin"
	
	"marketing/app/admin/apis"
	// "marketing/common/middleware/jwt"
)

func init() {
	allRouter = append(allRouter, registerSysRoleRouter)
}

func registerSysRoleRouter(r *gin.RouterGroup) {
	user := r.Group("/role")
	// user.Use(jwt.JWTAuth())
	{
		user.GET("/router", apis.GetRouter)
	}
}
