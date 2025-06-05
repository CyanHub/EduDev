package middleware

import (
	"ServerFramework/model/response"
	"ServerFramework/utils"
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
		j := utils.NewJwt()
		_, err := j.ParseToken(token)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
		}
	}
}
