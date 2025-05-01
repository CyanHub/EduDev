package middleware

import (
	"strconv"

	"github.com/CyanHub/EduDev/model/response"
	"github.com/CyanHub/EduDev/service"
	"github.com/CyanHub/EduDev/utils"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		path := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.RoleId))
		c.Set("authorityId", waitUse.RoleId)
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
