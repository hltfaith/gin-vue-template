package router

import (
	"github.com/gin-gonic/gin"
)

var allRouter = make([]func(r *gin.RouterGroup), 0)

func InitRouter(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	for _, f := range allRouter {
		f(api)
	}
	return r
}
