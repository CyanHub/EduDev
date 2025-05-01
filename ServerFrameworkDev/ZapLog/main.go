package main

import (
	_ "net/http/pprof"

	"ServerFramework/global"
	"ServerFramework/initialize"
)

func main() {
	initialize.MustConfig()
	initialize.MustLoadZap()
	initialize.MustInitDB()
	initialize.AutoMigrate(global.DB)
	//initialize.MustInitRedis()
	initialize.MustRunWindowServer()
}
