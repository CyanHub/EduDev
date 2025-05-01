package initialize

import (
	"IDPVerify/global"
	"IDPVerify/router"
	
	"fmt"

	"github.com/gin-gonic/gin"
)

// 创建`MustRunWindowServer`来实现`gin`服务的运行
func MustRunWindowServer() {
	engine := gin.Default()

	userGroup := router.UserGroup{}
	userGroup.InitUserRouters(engine)

	roleGroup := router.RoleGroup{}
	roleGroup.InitRoleRouters(engine)

	address := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	fmt.Println("启动服务器，监听端口：", address)
	if err := engine.Run(address); err != nil {
		panic(err)
	}

}

// package initialize

// import (
// 	"log"
// 	"github.com/gin-gonic/gin"
// )

// func MustRunWindowServer() {
// 	router := gin.Default()

// 	// Define your routes here
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello, World!",
// 		})
// 	})

// 	// Run the server
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatalf("Failed to run server: %v", err)
// 	}
// }
