package task

import (
	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/task"
)

func ClearOperationRecord(cronString string) {
	global.Cron.AddFunc(cronString, func() {
		task.ClearOperationRecord()
	})
}
