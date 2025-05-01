package global

import (
	"ServerFramework/config"

	"gorm.io/gorm"
)

// 在`global`目录下创建`global.go`文件，
// 并创建`config/Server`变量，以便后续在全局中使用配置文件
var (
	DB *gorm.DB
	CONFIG config.Config
	
)
