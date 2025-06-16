package request

type SetFilePermissions struct {
    FileID      uint64   `json:"fileId" binding:"required"`
    UserID      uint64   `json:"userId" binding:"required"`
    Permissions []string `json:"permissions" binding:"required"`
}

type FileUploadRequest struct {
    FileName string `json:"fileName" binding:"required"`
    FileType string `json:"fileType" binding:"required"`
    FileSize int64  `json:"fileSize" binding:"required"`

}

type FileDownloadRequest struct {
    FileID uint64 `json:"fileId" binding:"required"`
}
