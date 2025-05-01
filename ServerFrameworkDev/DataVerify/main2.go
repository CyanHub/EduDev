// package main

// import (
// 	"ServerFramework/config"
// 	"ServerFramework/global"
// 	"reflect"

// 	// "errors"

// 	// "ServerFramework/initialize"

// 	"fmt"
// 	"log"

// 	// "time"

// 	"github.com/fsnotify/fsnotify"
// 	"github.com/go-playground/locales/zh"
// 	ut "github.com/go-playground/universal-translator"
// 	"github.com/go-playground/validator/v10"
// 	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/spf13/viper"
// )

// func main() {

// 	// 数据校验 Week04
// 	ValidatorSimpleUse() // 数据校验简单使用

// 	// 服务器权限管理 Week03
// 	// Casbin的使用
// 	// GenerateToken() // 生成令牌
// 	// ParseToken(GenerateToken()) // 解析令牌
// 	// global.GetConfFromFlag() //

// 	// 身份认证JWT Week02
// 	// go BindEnv()
// 	// time.Sleep(time.Second*2)
// 	// initialize.MustLoadGorm()
// 	// initialize.MustRunWindowServer()
// 	// initialize.AutoMigrate(global.DB)
// 	// initialize.MustConfig()

// }

// // 数据校验 Week04

// var Validate *validator.Validate

// var zhTranslator ut.Translator
// var chTranslator ut.Translator

// type MainUser struct {
// 	Id       		int    `validate:"required,gte=1" chinese:"编号"`
// 	Username 		string `validate:"required,min=3,max=10,usernameUnique" chinese:"用户名"`
// 	// Username        string `validate:"required,min=3,max=10"`
// 	Name            string `validate:"required,min=3,max=10" chinese:"姓名"`
// 	Age             int    `validate:"required,gte=18,lte=100" chinese:"年龄"`
// 	Email           string `validate:"required,email" chinese:"邮箱"`
// 	Phone           string `validate:"e164" chinese:"手机号"`
// 	Password        string `validate:"required,min=8,max=16" chinese:"密码"`
// 	ConfirmPassword string `validate:"required,eqfield=Password" chinese:"确认密码"`
// 	Role            string `validate:"required,oneof=admin user" chinese:"角色"`
// 	Avatar          string `validate:"url" chinese:"头像"`
// 	LastLoginIp     string `validate:"ip" chinese:"最后登录IP"`
// 	// LastLoginIp     string `validate:"ip"`
// }

// var mainUsers = []MainUser{
// 	{
// 		Id:              1,
// 		Username:        "2201",
// 		Name:            "Alice",
// 		Age:             18,
// 		Email:           "12345678@123.com",
// 		Phone:           "+8612345678901",
// 		Password:        "12334566",
// 		ConfirmPassword: "12334566",
// 		Role:            "admin",
// 		Avatar:          "https://www.google.com",
// 		LastLoginIp:     "192.168.0.100",
// 	},
// }

// func ValidateUsernameUnique(fieldLevel validator.FieldLevel) bool {
// 	// 校验逻辑
// 	username := fieldLevel.Field().Interface().(string)

// 	// 查询数据库
// 	for _, user := range mainUsers {
// 		if user.Username == username {
// 			return false
// 		}
// 	}
// 	return true
// }

// func ValidateUsernameUnique2(fieldLevel validator.FieldLevel) bool {
// 	// 校验逻辑 其他逻辑
// 	username := fieldLevel.Field().Interface().(string)

// 	// 查询数据库
// 	for _, user := range mainUsers {
// 		if user.Username == username {
// 			return false
// 		}
// 	}
// 	return true
// }
// func ValidatorSimpleUse() {
// 	validate := validator.New()

// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		name := fld.Tag.Get("chinese")
// 		// 如果tag为-，则不返回字段名
// 		if name == "-" {
// 			return ""
// 		}
// 		return name
// 	})

// 	// 初始化中文翻译器
// 	zh := zh.New()
// 	uni := ut.New(zh, zh)

// 	// 获取中文翻译器
// 	zhTranslator, found := uni.GetTranslator("zh")
// 	if !found {
// 		panic("未找到中文翻译器 zh translator not found")
// 	}

// 	// 注册翻译，一般用来翻译自定义的校验器
// 	err := validate.RegisterTranslation("usernameUnique", zhTranslator, func(ut ut.Translator) error {
// 		return ut.Add("usernameUnique", "{0} {1} 已存在", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, err := ut.T("usernameUnique", fe.Field(), fe.Value().(string))
// 		if err != nil {
// 			panic(err)
// 		}
// 		return t
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = validate.RegisterValidation("usernameUnique", ValidateUsernameUnique)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 注册默认翻译
// 	err = zhTranslations.RegisterDefaultTranslations(validate, zhTranslator)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 在 TranslatorSimpleUse 函数中使用相同的翻译器
// 	zhTranslatorSimpleUse = zhTranslator

// 	onUser := MainUser{
// 		Id:              1,
// 		Username:        "2201",
// 		Name:            "Alice",
// 		Age:             18,
// 		Email:           "33061978@163.com",
// 		Phone:           "+8612345678901",
// 		Password:        "12345678",
// 		ConfirmPassword: "12345678",
// 		Role:            "admin",
// 		Avatar:          "https://www.baidu.com",
// 		LastLoginIp:     "192.168.0.220",
// 	}

// 	err = validate.Struct(onUser)
// 	if err != nil {
// 		fmt.Println(TranslatorSimpleUse(err))
// 	} else {
// 		fmt.Println("校验通过 Validate Success")
// 	}
// }

// var zhTranslatorSimpleUse ut.Translator

// func TranslatorSimpleUse(err error) string {
// 	errors := err.(validator.ValidationErrors)

// 	msg := ""
// 	for _, verr := range errors {
// 		msg += verr.Translate(zhTranslatorSimpleUse) + "\n"
// 	}
// 	return msg
// }

// // func ValidateUsernameUnique(fieldLevel validator.FieldLevel) bool {
// // 	// 校验逻辑
// // 	username := fieldLevel.Field().String()

// // 	// 查询数据库
// // 	for _, user := range mainUsers {
// // 		if user.Username == username {
// // 			return false
// // 		}
// // 	}
// // 	return true

// // }

// // ==================================================================================================
// // 服务器权限管理 Week03
// type CustomClaims struct {
// 	jwt.RegisteredClaims
// 	UserId     string
// 	Username   string
// 	Role       string
// 	Permission string
// }

// func GenerateToken() string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
// 		UserId:     "1",
// 		Username:   "11",
// 		Role:       "admin",
// 		Permission: "write",
// 	})
// 	fmt.Println(token.Valid)
// 	fmt.Println(token.Method)
// 	fmt.Println(token.Header)
// 	fmt.Println(token.Raw)
// 	fmt.Println(token.Claims)
// 	fmt.Println("=========================================================")
// 	tokenString, err := token.SignedString([]byte("2201"))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return tokenString
// }

// func ParseToken(tokenString string) {
// 	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		fmt.Println("内部token.Valid", token.Valid)
// 		return []byte("2201"), nil
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(token.Valid)
// 	fmt.Println(token.Raw)
// 	fmt.Println(token.Header)
// 	fmt.Println(token.Claims)
// }

// // =====================================================================================================
// func BindEnv() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalln(err)
// 	}
// 	viper.AutomaticEnv()
// 	if err := viper.BindEnv("mysql.password", "GoLand"); err != nil {
// 		log.Fatalln(err)
// 	}
// 	con := config.Config{}
// 	if err := viper.Unmarshal(&con); err != nil {
// 		log.Fatalln(err)
// 	}
// 	global.CONFIG = con

// 	viper.WatchConfig()
// 	viper.OnConfigChange(func(in fsnotify.Event) {
// 		if err := viper.Unmarshal(&con); err != nil {
// 			log.Fatalln(err)
// 		}
// 		global.CONFIG = con
// 		fmt.Println(con)
// 	})
// 	select {}
// }

// // import (
// // 	"ServerLearning/global"
// // 	"ServerLearning/initialize"
// // 	"fmt"

// // 	// "go/token"
// // 	// "log"
// // 	"time"

// // 	"github.com/golang-jwt/jwt/v5"
// // )

// // type CustonClainms struct {
// // 	jwt.RegisteredClaims
// // 	UserId     string
// // 	UserName   string
// // 	Roles      string
// // 	Permission string
// // }

// // // 定义 生成令牌 方法
// // func GenerateToken() string {
// // 	// V4 版本
// // 	// token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// // 	// 	"Subject":   "JWT 简单使用",
// // 	// 	"ExpiresAt": jwt.NewNumericDate(time.Now().Add(60 * time.Second)),
// // 	// })

// // 	// V5版本直接忽略最开始定义的token, _ 或者是 token, err,取而代之是直接到token
// // 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// // 		"Subject":   "JWT 简单使用",
// // 		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(60 * time.Second)),
// // 	})

// // 	fmt.Println(token.Valid)
// // 	fmt.Println(token.Method)
// // 	fmt.Println(token.Header)
// // 	fmt.Println(token.Raw)
// // 	fmt.Println(token.Claims)
// // 	fmt.Println("==================================================")

// // 	tokenString, _ := token.SignedString([]byte("2201"))

// // 	// if err != nil {
// // 	// 	// log.Fatalln(err)
// // 	// 	panic(err)
// // 	// }
// // 	fmt.Println(tokenString)
// // 	return tokenString
// // }

// // // 解析 生成令牌 方法
// // func ParseToken(tokenString string) {
// // 	token, _ := jwt.ParseWithClaims(tokenString, &CustonClainms{}, func(token *jwt.Token) (interface{}, error) {
// // 		return []byte("2201"), nil
// // 	})
// // 	fmt.Println(token.Valid)
// // 	fmt.Println(token.Method)
// // 	fmt.Println(token.Header)
// // 	fmt.Println(token.Raw)
// // 	fmt.Println(token.Claims)
// // }


package main