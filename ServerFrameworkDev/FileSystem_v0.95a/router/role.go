package router

import (
	"FileSystem/api"

	"github.com/gin-gonic/gin"
)

type RoleGroup struct{}

func (r *RoleGroup) InitRoleRouters(engine *gin.Engine) {
	roleRouters := engine.Group("role")
	roleRouters.POST("list", api.RoleList)
	roleRouters.POST("", api.CreateRole)
}


// package router

// import (
// 	"FileSystem/api"
// 	"FileSystem/middleware"

// 	"github.com/gin-gonic/gin"
// )

// type RoleGroup struct{}

// func (r *RoleGroup) InitRoleRouters(engine *gin.Engine) {
// 	roleRouters := engine.Group("role")
// 	roleRouters.POST("list", middleware.OperationRecord(), middleware.JwtMiddleware(), api.RoleList)
// 	roleRouters.POST("", middleware.OperationRecord(), middleware.JwtMiddleware(), api.RoleCreate)
// }
