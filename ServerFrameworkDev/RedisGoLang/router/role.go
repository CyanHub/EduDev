package router

import (
	"ServerFramework/api"

	"github.com/gin-gonic/gin"
)

type RoleGroup struct{}

func (r *RoleGroup) InitRoleRouters(engine *gin.Engine) {
	roleRouters := engine.Group("role")
	roleRouters.POST("list", api.RoleList)
	roleRouters.POST("", api.RoleCreate)
}


// package router

// import (
// 	"ServerFramework/api"
// 	"ServerFramework/middleware"

// 	"github.com/gin-gonic/gin"
// )

// type RoleGroup struct{}

// func (r *RoleGroup) InitRoleRouters(engine *gin.Engine) {
// 	roleRouters := engine.Group("role")
// 	roleRouters.POST("list", middleware.OperationRecord(), middleware.JwtMiddleware(), api.RoleList)
// 	roleRouters.POST("", middleware.OperationRecord(), middleware.JwtMiddleware(), api.RoleCreate)
// }
