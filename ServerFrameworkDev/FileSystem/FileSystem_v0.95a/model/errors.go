package model

import "gorm.io/gorm"

// RecordNotFound 封装 gorm 的 ErrRecordNotFound 错误
var RecordNotFound = gorm.ErrRecordNotFound
