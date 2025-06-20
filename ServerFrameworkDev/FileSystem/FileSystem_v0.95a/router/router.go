package router

import (
	"FileSystem/api"
	"FileSystem/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件服务
	r.Static("/static", "./static")
	r.Static("/css", "./static/css")
	r.Static("/images", "./static/images")
	r.Static("/js", "./static/js")

	// 页面路由
	r.StaticFile("/", "./static/pages/index.html")
	r.StaticFile("/index.html", "./static/pages/index.html")
	r.StaticFile("/register.html", "./static/pages/register.html")
	r.StaticFile("/login.html", "./static/pages/login.html")
	r.StaticFile("/list.html", "./static/pages/list.html")
	r.StaticFile("/logout.html", "./static/pages/logout.html")

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
			// 确保添加了 /avatar 接口路由
			userGroup.GET("/avatar", api.GetUserAvatar) 
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
	}

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

	// 使用GET方法 添加以下路由（暂时冗余）
	// r.GET("/register", func(c *gin.Context) {
	// 	c.File("./pages/register.html")
	// })

	return r
}
