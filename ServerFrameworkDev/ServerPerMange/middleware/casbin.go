package middleware

import (
	"ServerFramework/model/response"
	"ServerFramework/service"
	"ServerFramework/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Enforcer struct {
	
}
func CasbinMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        waitUse, _ := utils.GetClaims(c)
        // 获取请求的PATH
        path := c.Request.URL.Path
        // 获取请求方法
        act := c.Request.Method
        // 获取用户的角色
        sub := strconv.Itoa(int(waitUse.UserId))
        c.Set("authorityId", waitUse.Username)
        e := service.CasbinServiceApp.LoadCasbin()
        success, _ := e.Enforce(sub, path, act)
        if !success {
            response.FailWithDetailed(gin.H{}, "权限不足", c)
            c.Abort()
            return
        }
        c.Next()
    }
}