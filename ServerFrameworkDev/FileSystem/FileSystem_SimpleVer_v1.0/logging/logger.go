package logging

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func Init() {
	logPath := viper.GetString("logging.path") + "/app.log"
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapcore.Level(viper.GetInt("logging.level")),
	)

	Logger = zap.New(core, zap.AddCaller())
}

func LogInfo(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func LogError(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func LogAccess(userID uint, action string) {
	Logger.Info("user action",
		zap.Uint("user_id", userID),
		zap.String("action", action),
	)
}
