package initialize

import "ServerFramework/service"

// MustCasbin 初始化 Casbin 权限控制
// 该函数用于初始化 Casbin 权限控制，通过调用 LoadCasbin 方法来加载 Casbin 配置。
func MustCasbin() {
	// 创建 CasbinService 实例
	casbinService := &service.CasbinService{}
	// 通过实例指针调用 LoadCasbin 方法
	casbinService.LoadCasbin()
}
