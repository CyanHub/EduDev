package initialize

import (
	"fmt"
	ccore "github.com/CyanHub/EduDev/core"
	"github.com/CyanHub/EduDev/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustLoadZap() {
	levels := Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := ccore.NewZapCore(levels[i])
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
