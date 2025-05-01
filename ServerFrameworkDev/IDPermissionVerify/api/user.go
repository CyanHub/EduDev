package api

import (
	"IDPVerify/global"
	"IDPVerify/model"
	"IDPVerify/model/request"
	"IDPVerify/model/response"
	"IDPVerify/utils"
	"IDPVerify/service"

	"errors"
	"log"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	user, err := service.UserServiceApp.Login(req)
	if err != nil {
		if errors.Is(err, global.ErrUserNotFound) || errors.Is(err, global.ErrPasswordIncorrect) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println("登录失败: ", err)
			response.FailWithMessage("登录失败", c)
			return
		}
	}
	// userjwt:=utils.Jwt{}
	// var userTokenValue model.GoShopClaims
	// userTokenValue.BaseClaims.UserId=1
	// userTokenValue.BaseClaims.Username = user.Username
	// userTokenValue.ExpiresAt=jwt.NewNumericDate(time.Now().Add(time.Second*200))
	// usertoken,err:=userjwt.GenerateToken(&userTokenValue)
	if err != nil {
		log.Fatalf(err.Error())
	}


	response.OkWithData(user,c)

	jwt := utils.NewJwt()
	baseClaims := jwt.CreateClaims(model.BaseClaims{
		Username: "JWT",
		UserId:   1,
	})
	token, err := jwt.GenerateToken(&baseClaims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	successInfo := make(map[string]interface{})
	successInfo["user"] = user
	successInfo["token"] = token

	response.OkWithData(successInfo, c)
}

func Register(c *gin.Context) {
	var req request.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	user, err := service.UserServiceApp.Register(req)
	if err != nil {
		if errors.Is(err, global.ErrUserAlreadyExists) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println("注册失败: ", err)
			response.FailWithMessage("注册失败", c)
			return
		}
	}
	response.OkWithData(user, c)
}

func UserList(c *gin.Context) {
	var req request.UserListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	total, users, err := service.UserServiceApp.UserList(req)
	if err != nil {
		log.Println("获取用户列表失败: ", err)
		response.FailWithMessage("获取用户列表失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		Total:    total,
		List:     users,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}
