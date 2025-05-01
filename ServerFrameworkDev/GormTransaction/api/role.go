package api

import (
	"ServerFramework/global"
	"ServerFramework/model/request"
	"ServerFramework/model/response"
	"ServerFramework/service"

	"errors"
	"log"
	
	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	var req request.RoleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	total, list, err := service.RoleServiceApp.RoleList(req)
	if err != nil {
		log.Println("获取角色列表失败: ", err)
		response.FailWithMessage("获取角色列表失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		Total:    total,
		List:     list,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

func RoleCreate(c *gin.Context) {
	var req request.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	err := service.RoleServiceApp.RoleCreate(req)
	if err != nil {
		if errors.Is(err, global.ErrRoleAlreadyExists) {
			response.FailWithMessage("角色已存在", c)
			return
		} else {
			log.Println("创建角色失败: ", err)
			response.FailWithMessage("创建角色失败", c)
			return
		}
	}
	response.OkWithMessage("创建角色成功", c)
}
