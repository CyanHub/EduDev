package middleware

import (
	"strconv"
	"time"

	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/model/response"
	"FileSystem/service"
	"FileSystem/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户声明信息（已存在实现）
		waitUse, _ := utils.GetClaims(c)

		// 优化点：添加错误处理
		if waitUse == nil {
			global.Logger.Error("获取用户声明失败")
			response.FailWithCode(response.ERROR_TOKEN_INVALID, c)
			c.Abort()
			return
		}

		// 现有参数获取逻辑（保持原样）
		path := c.Request.URL.Path
		act := c.Request.Method
		sub := strconv.Itoa(int(waitUse.RoleID))

		// 优化点：添加调试日志
		// 在原有日志基础上增加详细上下文
		global.Logger.Debug("权限验证详情",
			zap.String("user", waitUse.Username),
			zap.Uint64("role", waitUse.RoleID),
			zap.String("path", path),
			zap.String("method", act))

		// 保持现有Casbin加载逻辑
		e := service.CasbinServiceApp.LoadCasbin()

		start := time.Now()
		defer func() {
			latency := time.Since(start)
			success := true
			global.Logger.Info("权限验证耗时",
				zap.Duration("latency", latency),
				zap.Bool("success", success))
		}()

		// 新增：带超时的权限检查
		done := make(chan bool)
		var success bool
		var err error

		go func() {
			success, err = e.Enforce(sub, path, act)
			done <- true
		}()

		select {
		case <-done:
		case <-time.After(1 * time.Second):
			global.Logger.Error("权限检查超时")
			response.FailWithMessage("权限验证超时", c)
			c.Abort()
			return
		}

		// 保持原有错误处理逻辑
		if err != nil {
			global.Logger.Error("权限检查出错", zap.Error(err))
			response.FailWithCode(403, c)
			// 记录操作
			record := model.OperationRecord{
				UserID:    uint64(waitUse.UserID),
				FileName:  path,
				Operation: act,
				Error:     "权限检查出错",
			}
			// 保存操作记录
			if err := service.OperationRecordServiceApp.CreateOperationRecord(record); err != nil {
				global.Logger.Error("记录操作失败", zap.Error(err))
			}
			c.Abort()
			return
		}
		if !success {
			response.FailWithCode(403, c)
			// 记录操作
			record := model.OperationRecord{
				UserID: uint64(waitUse.UserID),

				FileName:  path,
				Operation: act,
				Error:     "权限不足",
			}
			// 保存操作记录
			if err := service.OperationRecordServiceApp.CreateOperationRecord(record); err != nil {
				global.Logger.Error("记录操作失败", zap.Error(err))
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
