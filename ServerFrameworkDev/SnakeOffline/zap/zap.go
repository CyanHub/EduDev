package zap

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapWithCap() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	start := time.Now()
	time.Sleep(1 * time.Second)

	logger.Info("操作完成", zap.Duration("耗时", time.Since(start)))
}

func ZapWithStruct() {
	zap.NewExample()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	start := time.Now()
	time.Sleep(1 * time.Second)

	logger.Info("操作完成", zap.String("任务名称", "查询用户列表"), zap.Bool("查询成功", true), zap.Int("用户数量", 100), zap.String("耗时", time.Since(start).String()))
}

func ZapWithLevel() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("info level")
	logger.Debug("debug level")
	logger.Error("error level")
	logger.Warn("warn level")
}

func ZapWithSimpleFile() {
	// 打开 info.log 文件用于记录 Info 级别的日志
	file, err := os.OpenFile("./log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	core :=zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), zap.InfoLevel)
	logger := zap.New(core)
	defer logger.Sync()

	logger.Info("test")
}

func ZapWithFile() {

	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	cutter := NewCutter(
		"./logs",
		zap.InfoLevel,
		CutterWithLayout("2006-01-02 15-04"),
	)
	core := zapcore.NewCore(encoder, zapcore.AddSync(cutter), zap.InfoLevel)

	logger := zap.New(core, zap.AddCallerSkip(1))
	defer logger.Sync()

	logger.Info("test")
	logger.Error("error")
}

// func NewCore(level zapcore.Level, encoder zapcore.Encoder, writer zapcore.WriteSyncer) zapcore.Core {
// 	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
// 		return l == level
// 	})
// 	return zapcore.NewCore(encoder, writer, levelEnabler)
// }

// Cutter 实现 zapcore.WriteSyncer 接口
type Cutter struct {
	level        zapcore.Level        // 日志级别(debug, info, warn, error, dpanic, panic, fatal)
	layout       string        // 时间格式 2006-01-02 15:04:05
	director     string        // 日志文件夹
	file         *os.File      // 文件句柄
	mutex        *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}


func NewCutter(director string, level zapcore.Level, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:        level,
		director:     director,
		mutex:        new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	return rotate
}

// Write satisfies the io.Writer interface. It writes to the
// appropriate file handle that is currently being used.
// If we have reached rotation time, the target file gets
// automatically rotated, and also purged if necessary.
func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()

	values := make([]string, 0) 
	values = append(values, c.director)	
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout)+".log")
	}
	filename := filepath.Join(values...)
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}

func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}