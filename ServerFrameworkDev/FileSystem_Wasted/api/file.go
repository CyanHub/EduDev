package api

import (
	"FileSystem/global"
	"FileSystem/model/request"
	"FileSystem/model/response"
	"FileSystem/service"
	"FileSystem/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	// 处理JSON格式请求
	var req request.FileUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		global.Logger.Error("文件上传参数绑定失败", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	// 获取上传文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Error("获取上传文件失败", zap.Error(err))
		response.FailWithMessage("获取文件失败", c)
		return
	}
	defer file.Close()

	userId := utils.GetUserID(c)

	// 参数验证
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		global.Logger.Error("文件上传参数验证失败", zap.Error(err))
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	fileID, err := service.FileServiceApp.UploadFile(file, header, userId)
	if err != nil {
		global.Logger.Error("文件上传失败", zap.Error(err))
		response.FailWithMessage("文件上传失败", c)
		return
	}

	response.OkWithDetailed(gin.H{"fileID": fileID}, "上传成功", c)
}

func FileList(c *gin.Context) {
	userId := utils.GetUserID(c)
	files, err := service.FileServiceApp.GetFileList(userId)
	if err != nil {
		response.FailWithMessage("获取文件列表失败", c)
		return
	}
	response.OkWithData(files, c)
}

func DownloadFile(c *gin.Context) {
	var req request.FileDownloadRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	userID := utils.GetUserID(c)
	if ok := service.FileServiceApp.CheckFilePermission(userID, req.FileID, "download"); !ok {
		global.Logger.Warn("用户无文件下载权限",
			zap.Uint64("userID", userID),
			zap.Uint64("fileID", req.FileID))
		response.FailWithMessage("无文件下载权限", c)
		return
	}

	fileData, fileName, err := service.FileServiceApp.DownloadFile(req.FileID)
	if err != nil {
		global.Logger.Error("文件下载失败", zap.Error(err))
		response.FailWithMessage("文件下载失败", c)
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/octet-stream", fileData)
}

func SetFilePermissions(c *gin.Context) {
	var req request.SetFilePermissions
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	userID := utils.GetUserID(c)
	if ok := service.FileServiceApp.CheckFilePermission(userID, req.FileID, "set_permission"); !ok {
		global.Logger.Warn("用户无设置文件权限的权限",
			zap.Uint64("userID", userID),
			zap.Uint64("fileID", req.FileID))
		response.FailWithMessage("无设置文件权限的权限", c)
		return
	}

	if err := service.FileServiceApp.SetFilePermissions(req.FileID, req.UserID, req.Permissions); err != nil {
		global.Logger.Error("设置文件权限失败", zap.Error(err))
		response.FailWithMessage("设置文件权限失败", c)
		return
	}

	response.OkWithMessage("设置文件权限成功", c)
}

