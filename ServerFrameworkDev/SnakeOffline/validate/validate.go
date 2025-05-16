package validate

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var validate = validator.New()

type User struct {
	Id              int    `validate:"required,gte=1" chinese:"编号"`
	Username        string `validate:"required,min=3,max=10,usernameUnique" chinese:"用户名"`
	Name            string `validate:"required,min=3,max=10" chinese:"姓名"`
	Age             int    `validate:"required,gte=18,lte=100" chinese:"年龄"`
	Email           string `validate:"required,email" chinese:"邮箱"`
	Phone           string `validate:"e164" chinese:"手机号"`
	Password        string `validate:"required,min=8,max=16" chinese:"密码"`
	ConfirmPassword string `validate:"required,eqfield=Password" chinese:"确认密码"`
	Role            string `validate:"required,oneof=admin user" chinese:"角色"`
	Avatar          string `validate:"url" chinese:"头像"`
	LastLoginIp     string `validate:"ip" chinese:"最后登录ip"`
}

var users = []User{
	{ // 正常
		Id:              1,
		Username:        "1900300311",
		Name:            "John Doe",
		Age:             25,
		Email:           "john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 名字过长
		Id:              2,
		Username:        "1900300312",
		Name:            "Jane Doe With Long Name",
		Age:             25,
		Email:           "john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 邮箱格式错误
		Id:              3,
		Username:        "1900300313",
		Name:            "John Doe",
		Age:             25,
		Email:           "john.doeexample.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 年龄过大
		Id:              4,
		Username:        "1900300314",
		Name:            "John Doe",
		Age:             1000,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 手机号格式错误
		Id:              5,
		Username:        "1900300315",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "110120",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 密码过长
		Id:              6,
		Username:        "1900300316",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123withverylongpassword",
		ConfirmPassword: "password123withverylongpassword",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 确认密码错误
		Id:              7,
		Username:        "1900300317",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password1234",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 角色错误
		Id:              8,
		Username:        "1900300318",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "super_admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 头像格式错误
		Id:              9,
		Username:        "1900300319",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	},
	{ // 登录ip格式错误
		Id:              10,
		Username:        "1900300320",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "1111111111",
	},
}

func ValidateDemo() {
	for _, user := range users {
		err := validate.Struct(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("validate success")
		}
	}
}

func CustomValidate() {
	registerValidator()
	// 用户名和id都相同，自己本身，校验成功
	var user = User{
		Id:              10,
		Username:        "1900300320",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("validate success")
	}
	// 用户名和id不相同，校验失败
	var user2 = User{
		Id:              15,
		Username:        "1900300320",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password123",
		ConfirmPassword: "password123",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1.1",
	}
	err = validate.Struct(user2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("validate success")
	}
}

func verifyUsernameUnique(fl validator.FieldLevel) bool {
	// 获取字段值
	value := fl.Field().Interface().(string)
	// Id的值
	id := fl.Parent().FieldByName("Id").Interface().(int)
	// 校验是否重复
	for _, user := range users {
		if user.Id != id && user.Username == value {
			return false
		}
	}
	return true
}

func verifyWithTranslate() {
	// 注册自定义校验器
	registerValidator()
	// 注册翻译
	defineTranslator()
	var user = User{
		Id:              15,
		Username:        "1900300320",
		Name:            "John Doe",
		Age:             25,
		Email:           "+john.doe@example.com",
		Phone:           "+8613212345678",
		Password:        "password1234",
		ConfirmPassword: "password1234",
		Role:            "admin",
		Avatar:          "https://example.com/avatar.jpg",
		LastLoginIp:     "192.168.1",
	}
	err := validate.Struct(user)
	// 翻译错误消息
	fmt.Println(Translate(err))
}

func registerValidator() {
	err := validate.RegisterValidation("usernameUnique", verifyUsernameUnique)
	if err != nil {
		panic(err)
	}
}

var translator ut.Translator

func defineTranslator() {
	universalTranslator := ut.New(zh.New())
	translator, _ = universalTranslator.GetTranslator("zh")
	// 注册默认的翻译：就是将普通的英文转换为中文，不包括Key
	err := zhTranslations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		panic(err)
	}
	// 把字段转换为中文
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("chinese")
		// 如果tag为-，则不返回字段名
		if name == "-" {
			return ""
		}
		return name
	})
	// 注册翻译，一般用来翻译自定义的校验器
	validate.RegisterTranslation("usernameUnique", translator, func(ut ut.Translator) error {
		return ut.Add("usernameUnique", "{0} {1} 已存在", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("usernameUnique", fe.Field(), fe.Value().(string))
		return t
	})
	// validate.RegisterTranslation("password", translator, func(ut ut.Translator) error {
	// 	return ut.Add("password", "{0} 密码错误", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("password", fe.Field())
	// 	return t
	// })
}

func Translate(err error) string {
	// 仅翻译验证消息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return ""
	}
	msg := ""
	for _, err := range errs {
		msg += err.Translate(translator) + "\n"
	}
	return msg
}
