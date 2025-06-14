package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 在常量定义部分添加
const (
    ERROR   = 7
    SUCCESS = 0
    
    // 新增JWT相关错误码
    ERROR_TOKEN_NOT_EXIST = 1001
    ERROR_TOKEN_EXPIRE    = 1002
    ERROR_TOKEN_INVALID   = 1003
)

// 新增FailWithCode方法
func FailWithCode(code int, c *gin.Context) {
    var msg string
    switch code {
    case ERROR_TOKEN_NOT_EXIST:
        msg = "令牌不存在"
    case ERROR_TOKEN_EXPIRE:
        msg = "令牌已过期"
    case ERROR_TOKEN_INVALID:
        msg = "无效令牌"
    default:
        msg = "操作失败"
    }
    
    c.JSON(http.StatusOK, Response{
        Code: code,
        Data: nil,
        Msg:  msg,
    })
    c.Abort()
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// 注册成功后弹出提示框
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "注册成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		ERROR,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
