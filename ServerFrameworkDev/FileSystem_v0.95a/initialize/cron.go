package initialize

import (
	"FileSystem/global"
	"github.com/robfig/cron/v3"
)

func SetupCron() {
	global.Cron = cron.New(cron.WithSeconds())
	
	// 这里可以添加定时任务
	// 例如: task.AddClearOperationRecordTask(global.Cron)
	
	global.Cron.Start()
}
