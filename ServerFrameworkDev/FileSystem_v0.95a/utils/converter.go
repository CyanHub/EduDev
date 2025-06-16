package utils

import "strconv"

// 添加字符串转uint64方法
func StringToUint64(s string) uint64 {
	result, _ := strconv.ParseUint(s, 10, 64)
	return result
}
