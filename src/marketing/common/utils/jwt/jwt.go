package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"

	"marketing/common/models"
)

var jwtSecret = []byte("password")

type Claims struct {
	UserID   int64 `json:"user_id"`
	Username string `json:"username"`
	RoleID   uint `json:"role_id"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(u *models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	claims := Claims{
		u.UserID,
		u.Username,
		u.Rid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tag",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
