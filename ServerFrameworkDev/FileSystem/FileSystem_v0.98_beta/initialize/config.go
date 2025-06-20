package initialize

import (
	"flag"
	"fmt"
	"os"

	"FileSystem/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// MustConfig 初始化配置文件
// 该函数用于初始化配置文件，通过调用 viper.New 方法创建一个新的 viper.Viper 实例，然后设置配置文件的路径和类型，最后读取配置文件并将其解析到 global.CONFIG 结构体中。
// 如果命令行参数中指定了配置文件路径，则使用命令行参数中的路径，否则使用环境变量中的路径，最后使用默认的配置文件路径。
func MustConfig() {
	var config string
	flag.StringVar(&config, "c", "", "指定配置文件路径")
	flag.Parse()
	if config == "" { // 判断命令行参数是否指定了配置文件
		if configEnv := os.Getenv("CONFIG"); configEnv == "" { // 读取环境变量
			switch gin.Mode() {
			case gin.DebugMode:
				config = "config.yaml"
			case gin.TestMode:
				config = "config.test.yaml"
			case gin.ReleaseMode:
				config = "config.release.yaml"
			}
			fmt.Printf("您正在使用gin模式的%s环境名称, 配置文件的路径为%s\n", gin.Mode(), config)
		} else {
			config = configEnv
			fmt.Printf("您正在使用环境变量, 配置文件的路径为%s\n", configEnv)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值, 配置文件的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&global.CONFIG)
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变更: ", in.Name)
		err = v.Unmarshal(&global.CONFIG)
		if err != nil {
			panic(err)
		}
		fmt.Println(global.CONFIG)
	})
}
