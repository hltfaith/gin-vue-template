package router

import (
	"github.com/gin-gonic/gin"
	
	"marketing/app/admin/apis"
	"marketing/common/middleware/jwt"
)

func init() {
	allRouter = append(allRouter, registerSysUserRouter)
}

func registerSysUserRouter(r *gin.RouterGroup) {
	login := r.Group("")
	{
		login.POST("/login", apis.Login)
	}

	user := r.Group("/user")
	user.Use(jwt.JWTAuth())
	{
		user.GET("/info", apis.GetInfo)
		user.POST("/logout", apis.Logout)
	}
}
