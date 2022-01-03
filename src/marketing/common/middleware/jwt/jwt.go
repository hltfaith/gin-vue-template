package jwt

import (
	// "fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

	"marketing/common/utils/jwt"
	"marketing/common/utils/e"
)

// JWTAuth 中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 20000
		token := c.GetHeader("X-Token")
		if token == "" {
			code = e.ERROR
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"message": "Token信息错误",
			})
			c.Abort()
			return 
		}
		c.Next()
	}
}

