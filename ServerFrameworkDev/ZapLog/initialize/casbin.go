package initialize

import "github.com/CyanHub/EduDev/service"

func MustCasbin() {
	service.CasbinServiceApp.LoadCasbin()
}