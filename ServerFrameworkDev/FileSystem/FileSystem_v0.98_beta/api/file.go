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

// 添加上传接口（实验四的改进版）
func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Error("获取上传文件失败", zap.Error(err))
		response.FailWithMessage("获取文件失败", c)
		return
	}
	defer file.Close()

	userId := utils.GetUserID(c)

	// 添加请求参数绑定
	var req request.FileUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		global.Logger.Error("参数绑定失败", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}

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
    
    // 设置审计信息
    c.Set("audit_file", header.Filename)
    c.Set("audit_operation", "upload")
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
    
    // 设置审计信息
    c.Set("audit_file", fileName)
    c.Set("audit_operation", "download")
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

// 在现有接口基础上增加删除接口
func DeleteFile(c *gin.Context) {
    fileID := utils.StringToUint64(c.Param("id"))
    userID := utils.GetUserID(c)
    
    // 权限验证：仅允许文件所有者或管理员删除
    if !service.FileServiceApp.CheckFileOwner(userID, fileID) && 
       !service.UserServiceApp.IsAdmin(userID) {
        response.FailWithMessage("无删除权限", c)
        return
    }

    if err := service.FileServiceApp.DeleteFile(fileID); err != nil {
        global.Logger.Error("文件删除失败", zap.Error(err))
        response.FailWithMessage("文件删除失败", c)
        return
    }
    
    response.OkWithMessage("文件删除成功", c)
}
