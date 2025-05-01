package main

import (
	_ "net/http/pprof"

	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/initialize"
)

func main() {
	initialize.MustConfig()
	initialize.MustLoadZap()
	initialize.MustInitDB()
	initialize.AutoMigrate(global.DB)
	//initialize.MustInitRedis()
	initialize.MustRunWindowServer()
}
