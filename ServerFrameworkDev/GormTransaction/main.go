package main

import (
	"ServerFramework/config"
	"ServerFramework/global"
	"os"
	"time"

	// "ServerFramework/initialize"

	"fmt"
	"log"

	// "time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Week05
	// ZapWithLevel() // 日志级别使用
	// ZapSimpleUser() // 日志简单使用
	// ZapWithFile() // 日志文件使用
	// MustLoadZap() // 必须加载日志
	// 数据校验 Week04
	// ValidatorSimpleUseWith2201() // 数据校验简单使用

	// 服务器权限管理 Week03
	// Casbin的使用
	// GenerateToken() // 生成令牌
	// ParseToken(GenerateToken()) // 解析令牌
	// global.GetConfFromFlag() //

	// 身份认证JWT Week02
	// go BindEnv()
	// time.Sleep(time.Second*2)
	// initialize.MustLoadGorm()
	// initialize.MustRunWindowServer()
	// initialize.AutoMigrate(global.DB)
	// initialize.MustConfig()

}

// Gorm中的事务管理 Week07
// 01.使用事务的方式来实现数据的插入

// 服务器日志管理 Week05

// 00.zap对日志的简单使用
func ZapSimpleUser() {
	logger, err := zap.NewProduction()

	// zapcore.Level        // 级别
	// zapcore.LevelEnabler // 级别启用器
	// zapcore.WriteSyncer  // 写入器
	// zapcore.Encoder      // 编码器

	if err != nil {
		panic(err)
	}
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Error("Error")

	logger.Info("登录成功", zap.String("username", "202212070022"), zap.Int("age", 23), zap.Bool("status", true))
}

// 01.通过zap来实现日志等级区分
func ZapWithLevel() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("info level")
	logger.Debug("debug level")
	logger.Error("error level")
	logger.Warn("warn level")

	start := time.Time{}
	logger.Info("操作完成", zap.String("任务名称", "查询用户列表"), zap.Bool("查询成功", true), zap.Int("用户数量", 100), zap.String("耗时", time.Since(start).String()))
}

// 02.通过Zap来实现文件日志的输出
func ZapWithFile() {
	file, err := os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	systemLevel := zapcore.InfoLevel

	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == systemLevel
	})

	// defer file.Close()

	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), levelEnabler)

	logger := zap.New(core)

	// defer logger.Sync()

	logger = zap.New(core)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Error("Error")

	logger.Info("登录成功", zap.String("username", "202212070022"), zap.Int("age", 23), zap.Bool("status", true))
}

// 通过Zap来实现简单的文件日志的输出
func ZapWithSimpleFile() {
	file, err := os.OpenFile("./log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(file), zap.InfoLevel)
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("test")
}

// 03.通过Zap必须加载日志
// func MustLoadZap() {
// 	levels := Levels()
// 	length := len(levels)
// 	cores := make([]zapcore.Core, 0, length)
// 	for i := 0; i < length; i++ {
// 		core := ccore.NewZapCore(levels[i])
// 		cores = append(cores, core)
// 	}
// 	logger := zap.New(zapcore.NewTee(cores...))
// 	global.Logger = logger
// }

// func Levels() []zapcore.Level {
// 	levels := make([]zapcore.Level, 0, 7)
// 	level, err := zapcore.ParseLevel("info")
// 	if err != nil {
// 		level - zapcore.DebugLevel
// 	}
// 	for ; zapcore.Level; 		
// 	level <= zapcore.FatalLevel{
// 		levels = append(levels, level)}; level++ 

// 	}
// ==============================================================
// 数据校验 Week04
type MainUser struct {
	// Username string `validate:"required,min=3,max=10",username_unique`
	Id              int    `validate:"required,gte=1"`
	Username        string `validate:"required,min=3,max=10"`
	Name            string `validate:"required,min=3,max=10"`
	Age             int    `validate:"required,gte=18,lte=100"`
	Email           string `validate:"required,email"`
	Phone           string `validate:"e164"`
	Password        string `validate:"required,min=8,max=16"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
	Role            string `validate:"required,oneof=admin user"`
	Avatar          string `validate:"url"`
	LastLoginIp     string `validate:"ip"`
}

var mainUsers = []MainUser{
	{
		Id:              1,
		Username:        "2201",
		Name:            "Alice",
		Age:             18,
		Email:           "12345678@123.com",
		Phone:           "+8612345678901",
		Password:        "12334566",
		ConfirmPassword: "12334566",
		Role:            "admin",
		Avatar:          "https://www.google.com",
		LastLoginIp:     "192.168.0.100",
	},
}

func ValidateUsernameUnique(fieldLevel validator.FieldLevel) bool {
	// 校验逻辑
	username := fieldLevel.Field().String()

	// 查询数据库
	for _, user := range mainUsers {
		if user.Username == username {
			return false
		}
	}
	return true

}

func ValidatorSimpleUseWith2201() {
	validate := validator.New()

	err := validate.RegisterValidation("username_unique", ValidateUsernameUnique)
	if err != nil {
		panic(err)
	}
	universalTranslator := ut.New(zh.New())
	/*zhTranslator*/ _, found := universalTranslator.GetTranslator("zh")
	if !found {
		panic("未找到中文翻译器")
	}

	// zhTranslatio, _ := nil
	// err = zhTranslatio

	// onUser := MainUser{
	// 	Id:              1,
	// 	Username:        "2201",
	// 	Name:            "Alice",
	// 	Age:             18,
	// 	Email:           "33061978@163.com",
	// 	Phone:           "+8612345678901",
	// 	Password:        "12334566",
	// 	ConfirmPassword: "12334566",
	// 	Role:            "admin",
	// 	Avatar:          "https://www.baidu.com",
	// 	LastLoginIp:     "192.168.0.220",
	// }
	OnUser := MainUser{
		Id:              1,
		Username:        "2201",
		Name:            "Alice",
		Age:             18,
		Email:           "33061978@163.com",
		Phone:           "+8612345678901",
		Password:        "12334566",
		ConfirmPassword: "12334566",
		Role:            "admin",
		Avatar:          "https://www.baidu.com",
		LastLoginIp:     "192.168.0.220",
	}

	err = validate.Struct(OnUser)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("验证成功")
	}
}

// ==================================================================================================
// 服务器权限管理 Week03
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

// =====================================================================================================
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
