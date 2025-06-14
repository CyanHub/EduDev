package service

import (
	"FileSystem/global"
	"FileSystem/model"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type FileService struct{}

var FileServiceApp = new(FileService)

const (
	fileStorageRoot = "./storage/files"
	maxFileSize     = 100 << 20 // 100MB
)

func (f *FileService) UploadFile(file multipart.File, header *multipart.FileHeader, userID uint64) (uint64, error) {
	// 检查文件大小
	if header.Size > maxFileSize {
		return 0, errors.New("文件大小超过限制")
	}

	// 创建存储目录
	if err := os.MkdirAll(fileStorageRoot, 0755); err != nil {
		return 0, err
	}

	// 生成唯一文件名
	ext := filepath.Ext(header.Filename)
	newFilename := time.Now().Format("20060102150405") + ext
	filePath := filepath.Join(fileStorageRoot, newFilename)

	// 保存文件
	dst, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return 0, err
	}

	// 保存到数据库
	fileRecord := model.File{
		Name:       header.Filename,
		Path:       filePath,
		Size:       header.Size,
		Type:       ext,
		UserID:     userID,
		UploaderID: userID,
	}

	if err := global.DB.Create(&fileRecord).Error; err != nil {
		os.Remove(filePath) // 回滚文件
		return 0, err
	}

	return uint64(fileRecord.ID), nil
}

func (f *FileService) DownloadFile(fileID uint64) ([]byte, string, error) {
	var file model.File
	if err := global.DB.First(&file, fileID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("文件不存在")
		}
		return nil, "", err
	}

	content, err := os.ReadFile(file.Path)
	if err != nil {
		return nil, "", err
	}

	return content, file.Name, nil
}

func (f *FileService) CheckFilePermission(userID, fileID uint64, action string) bool {
	var file model.File
	if err := global.DB.First(&file, fileID).Error; err != nil {
		return false
	}

	// 上传者拥有所有权限
	if file.UploaderID == userID {
		return true
	}

	var permission model.FilePermission
	err := global.DB.Where("user_id = ? AND file_id = ?", userID, fileID).First(&permission).Error
	if err != nil {
		return false
	}

	switch action {
	case "download":
		return strings.Contains(permission.Permission, "read")
	case "delete":
		return strings.Contains(permission.Permission, "write")
	default:
		return false
	}
}

func (f *FileService) SetFilePermissions(fileID, userID uint64, permissions []string) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 清除旧权限
	if err := tx.Where("file_id = ? AND user_id = ?", fileID, userID).
		Delete(&model.FilePermission{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 添加新权限
	for _, perm := range permissions {
		if err := tx.Create(&model.FilePermission{
			FileID:     fileID,
			UserID:     userID,
			Permission: perm,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (f *FileService) GetFileList(userID uint64) ([]model.File, error) {
	var files []model.File
	err := global.DB.Where("user_id = ?", userID).Find(&files).Error
	return files, err
}
