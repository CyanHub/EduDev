package router

import (
	"FileSystem/api"
	"FileSystem/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件路由
	r.Static("/static", "./pages")
	r.Static("/images", "./images") // 新增头像资源路由
	r.StaticFile("/", "./pages/index.html")
	r.StaticFile("/login.html", "./pages/login.html")
	r.StaticFile("/register.html", "./pages/register.html")
	r.StaticFile("/logout.html", "./pages/logout.html") // 确保路径正确

	// API路由组
	apiGroup := r.Group("/api")
	{
		// 用户相关路由
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/login", api.Login)
			userGroup.POST("/register", api.Register)
			userGroup.POST("/logout", middleware.JwtMiddleware(), api.Logout)
		}

		// 文件相关路由（需要JWT验证）
		fileGroup := apiGroup.Group("/file").Use(middleware.JwtMiddleware())
		{
			fileGroup.POST("/upload", api.UploadFile)
			fileGroup.GET("/list", api.FileList)
			fileGroup.GET("/download/:id", api.DownloadFile)
			fileGroup.POST("/permission", api.SetFilePermissions)
		}
	}

	return r
}
