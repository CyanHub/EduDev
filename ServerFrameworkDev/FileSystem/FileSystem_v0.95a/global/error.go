package global

import "errors"

var (
	ErrNoToken      = errors.New("未提供认证token")
	ErrInvalidToken = errors.New("无效的token")
	// ...其他错误定义
)
