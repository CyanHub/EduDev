package main

import (
	"ConfigManage/config"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main1() {
	// GetEnv()
	// GetConfFromFlag()
	// ViperSimpleUse()
	DynamicLoad()
	// BindEnv()
}

// 获取环境变量参数
func GetEnv() {
	goland := os.Getenv("GoLand")
	if goland == "" {
		goland = "goland"
	}
	fmt.Println(goland)
}

// 使用flag获取参数配置
func GetConfFromFlag() {
	address := flag.String("add", "127.0.0.1", "server address")
	port := flag.Int("p", 8080, "server port")
	flag.Parse()
	fmt.Printf("server address is %s and server port is %d", *address, *port)
}

// viper库的简单使用
func ViperSimpleUse() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(viper.GetInt("mysql.port"))
	fmt.Println(viper.GetString("database"))
}

// 动态加载(热部署)
func DynamicLoad() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	conf := config.Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(conf)
	})
	select {}
}

// 在动态加载(热部署)前提下 绑定结构体
func BindEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	viper.AutomaticEnv()
	if err := viper.BindEnv("mysql.password", "Goland"); err != nil {
		log.Fatalln(err)
	}
	conf := config.Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&conf); err != nil {
			fmt.Println(conf)
		}
	})
	select {}
}