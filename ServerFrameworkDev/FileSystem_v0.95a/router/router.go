package router

import (
	"FileSystem/api"
	"FileSystem/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 新增CORS配置
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 添加静态文件路由
	// 修改静态文件路由配置
	r.Static("/static", "./static") // 将原有/static路由指向static目录
	r.Static("/uploads", "./uploads")

	// 新增头像资源路由（确保存在static/images/avatar目录）
	// 在静态文件路由部分添加
	r.Static("/static/images", "./static/images") // 头像资源路径
	// 新增css资源路由
	// r.Static("/css", "./static/css")

	// 添加审计中间件
	// r.Use(middleware.OperationRecord())

	// 修改页面路由配置
	r.StaticFile("/list.html", "./pages/list.html")
	r.StaticFile("/index.html", "./pages/index.html")
	r.StaticFile("/login.html", "./pages/login.html")
	r.StaticFile("/register.html", "./pages/register.html")
	r.StaticFile("/logout.html", "./pages/logout.html") // 确保路径正确

	// API路由组
	apiGroup := r.Group("/api")
	{
		// 用户相关路由
		// 在user路由组中添加信息接口
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/info", middleware.JwtMiddleware(), api.UserInfo)
			userGroup.POST("/login", api.Login)
			userGroup.POST("/register", api.Register) // 确保路由是/user/register
			userGroup.POST("/logout", middleware.JwtMiddleware(), api.Logout)
		}

		// 文件相关路由（需要JWT和Casbin验证）
		// 在file路由组补充缺失接口
		fileGroup := apiGroup.Group("/file").Use(middleware.JwtMiddleware(), middleware.CasbinMiddleware())
		{
			fileGroup.POST("/upload", api.UploadFile)
			fileGroup.GET("/list", api.FileList) // 确保该路由存在
			fileGroup.GET("/download/:id", api.DownloadFile)
			fileGroup.POST("/permission", api.SetFilePermissions)
			// 在file路由组中增加删除路由
			fileGroup.DELETE("/:id", api.DeleteFile)
		}

		// 角色权限管理（管理员专属）
		// 删除重复的authGroup分组，修改为：
		roleGroup := apiGroup.Group("/role").Use(middleware.AdminCheck())
		{
			roleGroup.POST("/roles", api.CreateRole)
			roleGroup.POST("/assign-permissions", api.AssignPermissions)
		}

		// 新增管理员路由组
		// adminGroup := apiGroup.Group("/admin").Use(middleware.AdminCheck())
		// {
		// 	adminGroup.POST("/roles", api.CreateRole)
		// 	adminGroup.POST("/assign-permissions", api.AssignPermissions)
		// }
	}


	return r
}
