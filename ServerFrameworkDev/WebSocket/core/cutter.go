package core

import (
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Cutter struct {
	level    zapcore.Level // 日志级别(debug, info, warn, error, dpanic, panic, fatal)
	layout   string        // 时间格式 2006-01-02 15:04:05
	director string        // 日志文件夹
	file     *os.File      // 文件句柄
	mutex    *sync.RWMutex // 读写锁
	zapcore.WriteSyncer
}

type CutterOption func(*Cutter)

func NewCutter(options ...CutterOption) *Cutter {
	cutter := &Cutter{
		level:    zapcore.InfoLevel,
		layout:   "2006-01-02 15:04:05",
		director: "./logs",
		mutex:    new(sync.RWMutex),
	}
	for _, option := range options {
		option(cutter)
	}
	return cutter
}

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

func CutterWithLevel(level zapcore.Level) CutterOption {
	return func(c *Cutter) {
		c.level = level
	}
}

func CutterWithDirector(director string) CutterOption {
	return func(c *Cutter) {
		c.director = director
	}
}

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
		values = append(values, time.Now().Format(c.layout), c.level.String()+".log")
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
