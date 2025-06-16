package global

import (
	"FileSystem/config"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	Cron   *cron.Cron
	Redis  *redis.Client
	Logger *zap.Logger
)
