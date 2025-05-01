package main

import (
	"RestfulAPI/user_server/config"
	"RestfulAPI/user_server/pkg/logs"
	"RestfulAPI/user_server/pkg/mysqldb"
	"RestfulAPI/user_server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	mysqldb.InitMysql()

	logs.InitLogger(
		config.CONFIG.Logger.LogTypes,
		config.CONFIG.Logger.Dir,
		logs.LogEnvType(config.CONFIG.System.Mode),
		config.CONFIG.Logger.LogMaxAge)

	gin.SetMode(config.CONFIG.System.Mode)

	r := gin.Default()

	r.SetTrustedProxies(nil)

	router.InitUserRouter(r)
	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}

}
