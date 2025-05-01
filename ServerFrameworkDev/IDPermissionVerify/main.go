package main

import (
	"IDPVerify/config"
	"IDPVerify/global"
	"IDPVerify/initialize"

	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// 身份认证JWT Week02

func main() {
	GenerateToken() // 生成令牌
	ParseToken(GenerateToken()) // 解析令牌
	global.GetConfFromFlag() // 

	go BindEnv()
	time.Sleep(time.Second*2)
	initialize.MustLoadGorm()
	initialize.MustRunWindowServer()
	initialize.AutoMigrate(global.DB)
	initialize.MustConfig()
	
}

// 
type CustomClaims struct {
	jwt.RegisteredClaims
	UserId     string
	Username   string
	Role       string
	Permission string
}

func GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserId:     "1",
		Username:   "11",
		Role:       "admin",
		Permission: "write",
	})
	fmt.Println(token.Valid)
	fmt.Println(token.Method)
	fmt.Println(token.Header)
	fmt.Println(token.Raw)
	fmt.Println(token.Claims)
	fmt.Println("=========================================================")
	tokenString, err := token.SignedString([]byte("2201"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenString
}

func ParseToken(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("内部token.Valid", token.Valid)
		return []byte("2201"), nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(token.Valid)
	fmt.Println(token.Raw)
	fmt.Println(token.Header)
	fmt.Println(token.Claims)
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
	global.CONFIG = con
	

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&con); err != nil {
			log.Fatalln(err)
		}
		global.CONFIG = con
		fmt.Println(con)
	})
	select {}
}

// import (
// 	"ServerLearning/global"
// 	"ServerLearning/initialize"
// 	"fmt"

// 	// "go/token"
// 	// "log"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// type CustonClainms struct {
// 	jwt.RegisteredClaims
// 	UserId     string
// 	UserName   string
// 	Roles      string
// 	Permission string
// }

// // 定义 生成令牌 方法
// func GenerateToken() string {
// 	// V4 版本
// 	// token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	// 	"Subject":   "JWT 简单使用",
// 	// 	"ExpiresAt": jwt.NewNumericDate(time.Now().Add(60 * time.Second)),
// 	// })

// 	// V5版本直接忽略最开始定义的token, _ 或者是 token, err,取而代之是直接到token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"Subject":   "JWT 简单使用",
// 		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(60 * time.Second)),
// 	})

// 	fmt.Println(token.Valid)
// 	fmt.Println(token.Method)
// 	fmt.Println(token.Header)
// 	fmt.Println(token.Raw)
// 	fmt.Println(token.Claims)
// 	fmt.Println("==================================================")

// 	tokenString, _ := token.SignedString([]byte("2201"))

// 	// if err != nil {
// 	// 	// log.Fatalln(err)
// 	// 	panic(err)
// 	// }
// 	fmt.Println(tokenString)
// 	return tokenString
// }

// // 解析 生成令牌 方法
// func ParseToken(tokenString string) {
// 	token, _ := jwt.ParseWithClaims(tokenString, &CustonClainms{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("2201"), nil
// 	})
// 	fmt.Println(token.Valid)
// 	fmt.Println(token.Method)
// 	fmt.Println(token.Header)
// 	fmt.Println(token.Raw)
// 	fmt.Println(token.Claims)
// }
