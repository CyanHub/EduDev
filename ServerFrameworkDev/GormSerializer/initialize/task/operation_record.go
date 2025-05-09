package task

import (
	"ServerFramework/global"
	"ServerFramework/task"
)

func ClearOperationRecord(cronString string) {
	global.Cron.AddFunc(cronString, func() {
		task.ClearOperationRecord()
	})
}
