package config

import "github.com/spf13/viper"

type Server struct {
	MySQL  MySQL
	Redis  Redis
	App    App
	Jwt    Jwt
	Logger Logger
}

// 需要添加环境变量覆盖支持
func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FS") // 环境变量前缀 FS_

	// 绑定环境变量映射
	viper.BindEnv("mysql.host", "FS_MYSQL_HOST")
	viper.BindEnv("mysql.port", "FS_MYSQL_PORT")
	viper.BindEnv("mysql.user", "FS_MYSQL_USER")
	viper.BindEnv("mysql.password", "FS_MYSQL_PASSWORD")
	viper.BindEnv("mysql.database", "FS_MYSQL_DATABASE")
	viper.BindEnv("mysql.config", "FS_MYSQL_CONFIG")
	viper.BindEnv("mysql.csv_sep", "FS_MYSQL_CSV_SEP")
}
