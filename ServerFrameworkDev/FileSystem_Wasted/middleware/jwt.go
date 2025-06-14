package middleware

import (
	"FileSystem/model/response"
	"FileSystem/utils"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取token，并验证token是否为空
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		// 2. 解析token
		j := utils.NewJWT()
		_, err := j.ParseToken(token)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
		}
	}
}

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            response.FailWithCode(response.ERROR_TOKEN_NOT_EXIST, c)
            c.Abort()
            return
        }

        claims, err := utils.NewJWT().ParseToken(token)
        if err != nil {
            if err == utils.ErrTokenExpired {
                response.FailWithCode(response.ERROR_TOKEN_EXPIRE, c)
                c.Abort()
                return
            }
            response.FailWithCode(response.ERROR_TOKEN_INVALID, c)
            c.Abort()
            return
        }

        c.Set("claims", claims)
        c.Next()
    }
}
