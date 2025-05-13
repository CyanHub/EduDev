package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Env string

const (
	ConfigEnvDebug   Env = "debug"
	ConfigEnvRelease Env = "release"
)

// mysql connection's info
type Mysql struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         string `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
}

// redisdb connection's info
type Redis struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
	Password string `mapstructure:"password" json:"password"`
}

// kafka connection's info
type Kafka struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

//type ElasticSearch struct {
//	Host string `mapstructure:"host" json:"host"`
//}

// the system's info
type System struct {
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

// tls
//type Tls struct {
//	Enable bool   `mapstructure:"enable" json:"enable"`
//	Cert   string `mapstructure:"cert" json:"cert"`
//	Key    string `mapstructure:"key" json:"key"`
//}

// logger
type Logger struct {
	Stdout    bool     `mapstructure:"stdout" json:"stdout"`
	Level     string   `mapstructure:"level" json:"level"`
	Dir       string   `mapstructure:"dir" json:"dir"`
	Rotation  bool     `mapstructure:"rotation" json:"rotation"`
	LogMaxAge int      `mapstructure:"logMaxAge" json:"logMaxAge"`
	LogTypes  []string `mapstructure:"logTypes" json:"logTypes"`
}

// all configs's info
type Config struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis"`
	System System `mapstructure:"system" json:"system"`
	Logger Logger `mapstructure:"logger" json:"logger"`
	Kafka  Kafka  `mapstructure:"kafka" json:"kafka"`
}

var (
	CONFIG Config
	VP     *viper.Viper
)

// the init
func Init() {
	// 获取当前工作目录
	workingDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("无法获取工作目录: %v", err))
	}
	// 拼接 config.yaml 文件的绝对路径
	configPath := fmt.Sprintf("%s/../config/config.yaml", workingDir)

	v := viper.New()
	v.SetConfigFile(configPath)
	err = v.ReadInConfig()
	if err != nil {
		// 待办事项符合统一的待办事项错误
		fmt.Println(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configs 文件已更改:", e.Name)
		if err = v.Unmarshal(&CONFIG); err != nil {
			// 待办事项符合统一的待办事项错误
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&CONFIG); err != nil {
		// 待办事项符合统一的待办事项错误
		fmt.Println(err)
	}
	VP = v
	fmt.Println(VP)
}
