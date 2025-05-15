package initialize

import "ServerFramework/service"

func MustCasbin() {
	// 创建 CasbinService 实例
	casbinService := &service.CasbinService{}
	// 通过实例指针调用 LoadCasbin 方法
	casbinService.LoadCasbin()
}
