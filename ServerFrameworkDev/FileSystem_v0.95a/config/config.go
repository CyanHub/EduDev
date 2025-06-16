package config

import (
	"github.com/spf13/viper"
)

// MySQL 数据库配置结构体
type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Config   string `json:"config"`
	CsvSep   string `json:"csv_sep"`
}

// Redis 配置结构体
type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// Jwt JWT 配置结构体
type Jwt struct {
	SigningKey  string `json:"signing_key"`
	ExpireTime  int    `json:"expire_time"`
	TokenPrefix string `json:"token_prefix"`
	Secret      string `json:"secret"`  // 添加 Secret 字段
	Issuer      string `json:"issuer"`  // 添加 Issuer 字段
}

// Logger 日志配置结构体
type Logger struct {
	Level      string `json:"level"`
	Path       string `json:"path"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
	Director   string `json:"director"`
	Layout     string `json:"layout"`
}

// App 应用配置结构体
type App struct {
	Env  string `json:"env"`
	Port int    `json:"port"`
}

// Server 服务器配置结构体，包含各个模块的配置
type Server struct {
	MySQL  MySQL  `json:"mysql"`
	Redis  Redis  `json:"redis"`
	App    App    `json:"app"`
	Jwt    Jwt    `json:"jwt"`
	Logger Logger `json:"logger"`
	Issuer string `json:"issuer"`
	Secret string `json:"secret"`
}

// 初始化环境变量绑定
func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FS") // 环境变量前缀 FS_

	// 绑定 MySQL 环境变量映射
	viper.BindEnv("mysql.host", "FS_MYSQL_HOST")
	viper.BindEnv("mysql.port", "FS_MYSQL_PORT")
	viper.BindEnv("mysql.user", "FS_MYSQL_USER")
	viper.BindEnv("mysql.password", "FS_MYSQL_PASSWORD")
	viper.BindEnv("mysql.database", "FS_MYSQL_DATABASE")
	viper.BindEnv("mysql.config", "FS_MYSQL_CONFIG")
	viper.BindEnv("mysql.csv_sep", "FS_MYSQL_CSV_SEP")

	// 绑定 Redis 环境变量映射
	viper.BindEnv("redis.addr", "FS_REDIS_ADDR")
	viper.BindEnv("redis.password", "FS_REDIS_PASSWORD")
	viper.BindEnv("redis.db", "FS_REDIS_DB")

	// 绑定 JWT 环境变量映射
	viper.BindEnv("jwt.signing_key", "FS_JWT_SIGNING_KEY")
	viper.BindEnv("jwt.expire_time", "FS_JWT_EXPIRE_TIME")
	viper.BindEnv("jwt.token_prefix", "FS_JWT_TOKEN_PREFIX")
	viper.BindEnv("jwt.secret", "FS_JWT_SECRET")  // 添加 Secret 字段绑定
	viper.BindEnv("jwt.issuer", "FS_JWT_ISSUER")  // 添加 Issuer 字段绑定

	// 绑定 Logger 环境变量映射
	viper.BindEnv("logger.level", "FS_LOGGER_LEVEL")
	viper.BindEnv("logger.path", "FS_LOGGER_PATH")
	viper.BindEnv("logger.max_size", "FS_LOGGER_MAX_SIZE")
	viper.BindEnv("logger.max_backups", "FS_LOGGER_MAX_BACKUPS")
	viper.BindEnv("logger.max_age", "FS_LOGGER_MAX_AGE")
	viper.BindEnv("logger.compress", "FS_LOGGER_COMPRESS")
	viper.BindEnv("logger.director", "FS_LOGGER_DIRECTOR")
	viper.BindEnv("logger.layout", "FS_LOGGER_LAYOUT")

	// 绑定 App 环境变量映射
	viper.BindEnv("app.env", "FS_APP_ENV")
}
