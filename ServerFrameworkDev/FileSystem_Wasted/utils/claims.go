package utils

import (
	"FileSystem/model"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	return token
}

func GetClaims(c *gin.Context) (*model.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetUserId(c *gin.Context) uint64 {
	claims, err := GetClaims(c)
	if err != nil {
		return 0
	}
	return claims.UserID
}

func GetUserName(c *gin.Context) string {
	claims, err := GetClaims(c)
	if err != nil {
		return ""
	}
	return claims.Username
}
