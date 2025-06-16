package router

import (
	"FileSystem/api"
	"FileSystem/middleware"

	"github.com/gin-gonic/gin"
)

type UserGroup struct{}

func (u *UserGroup) InitUserRouters(engine *gin.Engine) {
	userRouters := engine.Group("/api/user")
	{
		userRouters.POST("/login", api.Login)
		userRouters.POST("/register", api.Register)
		userRouters.POST("/logout", api.Logout)
		userRouters.POST("/list", middleware.OperationRecord(), middleware.JwtMiddleware(), middleware.CasbinMiddleware(), api.UserList)
		userRouters.GET("/avatar", api.GetUserAvatar)
	}
}
	