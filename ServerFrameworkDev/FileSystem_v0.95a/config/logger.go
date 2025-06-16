package config

// 修改日志配置
type Logger struct {
    Level      string `json:"level"`
    Path       string `json:"path"`  // 修改为具体路径 "./logs/system.log"
    MaxSize    int    `json:"maxSize"`
    MaxBackups int    `json:"maxBackups"`
    MaxAge     int    `json:"maxAge"`
    Compress   bool   `json:"compress"`
	Director string `json:"director"`
	Layout   string `json:"layout"`
}
