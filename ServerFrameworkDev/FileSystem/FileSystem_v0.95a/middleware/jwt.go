package middleware

import (
	"FileSystem/model/response"
	"FileSystem/utils"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithCode(response.ERROR_TOKEN_NOT_EXIST, c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.ErrTokenExpired {
				response.FailWithCode(response.ERROR_TOKEN_EXPIRE, c)
			} else {
				response.FailWithCode(response.ERROR_TOKEN_INVALID, c)
			}
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("claims", claims)
		c.Next()
	}
}

// 修改AdminCheck函数
func AdminCheck() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, exists := c.Get("claims")
        if !exists {
            response.FailWithCode(response.ERROR_TOKEN_INVALID, c)
            c.Abort()
            return
        }

        // 根据实际Claims结构调整
        userClaims, ok := claims.(*utils.Claims) 
        if !ok || userClaims.RoleID != 1 { // 假设1是管理员角色ID
            response.FailWithCode(response.ERROR_NO_ADMIN_PERMISSION, c)
            c.Abort()
            return
        }
        c.Next()
    }
}

// 移除冗余的JWTAuth方法，统一使用JwtMiddleware
