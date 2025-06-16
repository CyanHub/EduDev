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

type Claims struct {
    UserID   uint64
    Username string
    RoleID   uint64 // 确保存在角色ID字段
    // 其他必要字段...
}

// 添加角色获取方法
func (c *Claims) GetRoleID() uint64 {
    return c.RoleID
}
