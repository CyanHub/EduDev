package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/service"
	"FileSystem/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var operationRecordService = service.OperationRecordServiceApp

var (
	respPool   sync.Pool
	bufferSize = 1024
)

// 定义操作类型常量
const (
	OperationUpload   = "upload"
	OperationDownload = "download"
	OperationDelete   = "delete"
)

// type OperationRecord struct {
// 	Ip          string    `gorm:"size:128;comment:请求IP"`
// 	Method      string    `gorm:"size:16;comment:请求方法"`
// 	Path        string    `gorm:"size:255;comment:请求路径"`
// 	Agent       string    `gorm:"size:255;comment:代理信息"`
// 	Body        string    `gorm:"type:text;comment:请求体"`
// 	Resp        string    `gorm:"type:text;comment:响应体"`
// 	UserID      int       `gorm:"comment:用户ID"`
// 	Error       string    `gorm:"type:text;comment:错误信息"`
// 	Latency     time.Duration
// 	Status      int       `gorm:"comment:状态码"`
// 	FileName    string    `gorm:"size:255;comment:操作文件名"`
// 	Operation   string    `gorm:"size:32;comment:操作类型"`
// 	CreatedAt   time.Time `gorm:"comment:创建时间"`
// }

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int

		// 获取请求体
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Error("读取请求体失败", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		// 获取用户ID
		claims, _ := utils.GetClaims(c)
		if claims != nil && claims.BaseClaims.UserID != 0 {
			userId = int(claims.BaseClaims.UserID)
		} else {
			id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}

		// 创建审计记录
		record := model.OperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   "",
			UserID: uint64(userId),
		}

		// 处理文件上传类型的请求体
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			record.Body = "[文件内容]"
			record.Operation = OperationUpload
		} else {
			if len(body) > bufferSize {
				record.Body = "[超出记录长度]"
			} else {
				record.Body = string(body)
			}
		}

		// 捕获响应
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		startTime := time.Now()

		c.Next()

		// 记录操作结果
		latency := time.Since(startTime)
		record.Error = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		// 处理文件下载类型的响应
		if isFileDownload(c) {
			record.Operation = OperationDownload
			if len(record.Resp) > bufferSize {
				record.Resp = "[文件内容]"
			}
		}

		// 记录文件名（如果存在）
		if fileName, exists := c.Get("audit_file"); exists {
			record.FileName = fileName.(string)
		}

		// 保存审计记录
		if err := operationRecordService.CreateOperationRecord(record); err != nil {
			global.Logger.Error("创建操作记录失败", zap.Error(err))
		}
	}
}

func isFileDownload(c *gin.Context) bool {
	headers := []string{
		"application/force-download",
		"application/octet-stream",
		"application/vnd.ms-excel",
		"application/download",
	}

	for _, header := range headers {
		if strings.Contains(c.Writer.Header().Get("Content-Type"), header) {
			return true
		}
	}
	return strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment")
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
