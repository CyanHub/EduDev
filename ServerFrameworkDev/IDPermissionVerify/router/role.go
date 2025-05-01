package router

import (
	"IDPVerify/api"
	"IDPVerify/middleware"

	"github.com/gin-gonic/gin"
)

type RoleGroup struct {
}

func (r *RoleGroup) InitRoleRouters(engine *gin.Engine) {
	roleRouters := engine.Group("role")
	roleRouters.Use(middleware.JwtMiddleware()) // 该路由组的全局中间件
	roleRouters.POST("list", api.RoleList)
	roleRouters.POST("", api.RoleCreate)

}
