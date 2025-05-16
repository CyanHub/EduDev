package viper

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ReadEnv() {
	dbPassword := os.Getenv("DBPASSWORD")
	if dbPassword == "" {
		dbPassword = "123456"
	}
	fmt.Println(dbPassword)
}

func ReadFlag() {
	var port int
	flag.IntVar(&port, "port", 8080, "Server port")
	flag.Parse()
	fmt.Println("Server running on port:", port)
}

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}
	fmt.Println(viper.Get("db.password"))
}

func ReadConfigWithHotReload() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./viper")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	// 模拟持续运行
	select {}
}

func BindEnv() {
	viper.BindEnv("db.password", "DBPASSWORD")
	fmt.Println(viper.Get("db.password"))
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Config struct {
	DB DB `mapstructure:"db" yaml:"db"`
}

func BindStruct() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}
	var config Config
	viper.Unmarshal(&config)
	fmt.Println(config)
}
