package utils

import (
	"ServerFramework/model"

	"github.com/gin-gonic/gin"
)

// 集中处理与用户相关的操作，比如从 `JWT` 中提取用户信息（例如 `UserId` 和 `UserName`）等。
// 专门负责从请求中获取 `JWT`、解析 `claims`、提取用户信息等操作.

// GetToken`函数，从请求中获取`Authorization`请求头
func GetToken(c *gin.Context) string {
    token := c.Request.Header.Get("Authorization")
    return token
}

// `GetClaims`函数，从token中获取声明（claims）

func GetClaims(c *gin.Context) (*model.GoShopClaims, error) {
    token := GetToken(c)
    j := NewJwt()
    claims, err := j.ParseToken(token)
    if err != nil {
        return nil, err
    }
    return claims, nil
}

// GetUserId`函数，根据请求头中的token解析得到声明，并返回声明中的`UserId`
func GetUserId(c *gin.Context) uint64 {
    claims, err := GetClaims(c)
    if err != nil {
        return 0
    }
    return claims.UserId
}

// `GetUserName`函数，根据请求头中的token解析得到声明，并返回声明中的`UserName`
func GetUserName(c *gin.Context) string {
    claims, err := GetClaims(c)
    if err != nil {
       return ""
    }
    return claims.Username
}

// 