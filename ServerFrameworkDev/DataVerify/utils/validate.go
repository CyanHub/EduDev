package utils

import (
	"reflect"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Validate *validator.Validate

var zhTranslator ut.Translator
var chTranslator ut.Translator

type MainUser struct {
	Id       		int    `validate:"required,gte=1" chinese:"编号"`
	Username 		string `validate:"required,min=3,max=10,usernameUnique" chinese:"用户名"`
	// Username        string `validate:"required,min=3,max=10"`
	Name            string `validate:"required,min=3,max=10" chinese:"姓名"`
	Age             int    `validate:"required,gte=18,lte=100" chinese:"年龄"`
	Email           string `validate:"required,email" chinese:"邮箱"`
	Phone           string `validate:"e164" chinese:"手机号"`
	Password        string `validate:"required,min=8,max=16" chinese:"密码"`
	ConfirmPassword string `validate:"required,eqfield=Password" chinese:"确认密码"`
	Role            string `validate:"required,oneof=admin user" chinese:"角色"`
	Avatar          string `validate:"url" chinese:"头像"`
	LastLoginIp     string `validate:"ip" chinese:"最后登录IP"`
	// LastLoginIp     string `validate:"ip"`
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

// ValidateUsernameUnique 是一个自定义验证函数，用于验证用户名是否唯一。
// 该函数会遍历已存在的用户列表，检查传入的用户名是否已被使用。
// 参数 fieldLevel 是 validator.FieldLevel 类型，包含了当前验证字段的相关信息。
// 返回值为 bool 类型，若用户名唯一则返回 true，否则返回 false。
func ValidateUsernameUnique(fieldLevel validator.FieldLevel) bool {
	// 从 fieldLevel 中获取当前验证字段的值，并将其转换为 string 类型
	username := fieldLevel.Field().Interface().(string)

	// 模拟查询数据库，遍历已存在的用户列表 mainUsers
	for _, user := range mainUsers {
		// 检查当前用户名是否与列表中已存在的用户名相同
		if user.Username == username {
			// 若相同，则表示用户名不唯一，返回 false
			return false
		}
	}
	// 若遍历完整个列表都未找到相同的用户名，则表示用户名唯一，返回 true
	return true
}

func ValidatorSimpleUse() {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("chinese")
		// 如果tag为-，则不返回字段名
		if name == "-" {
			return ""
		}
		return name
	})

	// 初始化中文翻译器
	zh := zh.New()
	uni := ut.New(zh, zh)

	// 获取中文翻译器
	zhTranslator, found := uni.GetTranslator("zh")
	if !found {
		panic("未找到中文翻译器 zh translator not found")
	}

	// 注册翻译，一般用来翻译自定义的校验器
	err := validate.RegisterTranslation("usernameUnique", zhTranslator, func(ut ut.Translator) error {
		return ut.Add("usernameUnique", "{0} {1} 已存在", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T("usernameUnique", fe.Field(), fe.Value().(string))
		if err != nil {
			panic(err)
		}
		return t
	})
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("usernameUnique", ValidateUsernameUnique)
	if err != nil {
		panic(err)
	}

	// 注册默认翻译
	err = zhTranslations.RegisterDefaultTranslations(validate, zhTranslator)
	if err != nil {
		panic(err)
	}

	// 在 TranslatorSimpleUse 函数中使用相同的翻译器
	zhTranslatorSimpleUse = zhTranslator

	onUser := MainUser{
		Id:              1,
		Username:        "2201",
		Name:            "Alice",
		Age:             18,
		Email:           "33061978@163.com",
		Phone:           "+8612345678901",
		Password:        "12345678",
		ConfirmPassword: "12345678",
		Role:            "admin",
		Avatar:          "https://www.baidu.com",
		LastLoginIp:     "192.168.0.2200000",
	}

	err = validate.Struct(onUser)
	if err != nil {
		fmt.Println(TranslatorSimpleUse(err))
	} else {
		fmt.Println("校验通过 Validate Success")
	}
}

var zhTranslatorSimpleUse ut.Translator

// TranslatorSimpleUse 函数用于将验证错误信息翻译为中文。
// 它接收一个 error 类型的参数 err，该参数通常是由 validator 包校验结构体时返回的错误。
// 函数会将错误信息转换为中文并拼接成一个字符串返回。
func TranslatorSimpleUse(err error) string {
	// 将传入的错误类型断言为 validator.ValidationErrors 类型，
	// 该类型包含了多个验证错误信息。
	errors := err.(validator.ValidationErrors)

	// 初始化一个空字符串，用于存储拼接后的翻译错误信息。
	msg := ""
	// 遍历所有的验证错误信息
	for _, verr := range errors {
		// 调用 Translate 方法将单个验证错误信息翻译为中文，
		// 并添加换行符后拼接到 msg 字符串中。
		msg += verr.Translate(zhTranslatorSimpleUse) + "\n"
	}
	// 返回拼接好的翻译错误信息字符串
	return msg
}

