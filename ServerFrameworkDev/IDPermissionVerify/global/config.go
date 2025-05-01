package global

import (
	"IDPVerify/config"
	
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func GetEnv() {
	goland := os.Getenv("GoLand")
	if goland == "" {
		goland = "goland"
	}
	fmt.Println(goland)
}

func GetConfFromFlag() {
	add := flag.String("add", "127.0.0.1", "sever address")
	port := flag.Int("port", 8080, "server port")
	flag.Parse()
	fmt.Printf("server address is %s and server port is %d\n", *add, *port)
	CONFIG.Server.Port = *port

	fmt.Println(CONFIG.Server.Port)

}

func ViperSimpleUse() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(viper.GetInt("mysql.port"))

	fmt.Println(viper.GetString("database"))
}

func DynamicLoad() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	con := config.Config{}
	if err := viper.Unmarshal(&con); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(con)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&con); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(con)
	})
	select {}
}

func BindEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	viper.AutomaticEnv()
	if err := viper.BindEnv("mysql.password", "GoLand"); err != nil {
		log.Fatalln(err)
	}
	con := config.Config{}
	if err := viper.Unmarshal(&con); err != nil {
		log.Fatalln(err)
	}
	CONFIG = con
	

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&con); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(con)
	})
	select {}
}
