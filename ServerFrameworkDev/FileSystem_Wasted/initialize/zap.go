package initialize

import (
	"fmt"

	"FileSystem/core"
	"FileSystem/global"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MustLoadZap 初始化 zap 日志库
// 该函数用于初始化 zap 日志库，通过调用 zap.New 方法创建一个新的 zap.Logger 实例。
func MustLoadZap() {
	levels := Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := core.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	global.Logger = logger
}

// Levels 根据字符串转化为 zapcore.Levels
func Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(global.CONFIG.Logger.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	fmt.Println(levels)
	return levels
}

func InitLogger() {
	// 使用 Viper 加载日志配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}

	var loggerConfig struct {
		Logger struct {
			Level    string `yaml:"level"`
			Director string `yaml:"director"`
			Layout   string `yaml:"layout"`
		} `yaml:"logger"`
	}

	if err := viper.Unmarshal(&loggerConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// 配置日志切割
	writeSyncer := getLogWriter(
		loggerConfig.Logger.Director,
		100,
		30,
		7,
	)

	// 配置日志级别
	level := getLogLevel(loggerConfig.Logger.Level)

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writeSyncer,
		level,
	)

	// 创建日志记录器
	global.Logger = zap.New(core, zap.AddCaller())

	// 记录初始化信息
	global.Logger.Info("Logger initialized successfully")
}

// getLogWriter 配置日志切割
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getLogLevel 解析日志级别
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
