package task

import (
	"ServerFramework/global"
	"ServerFramework/model"
	"ServerFramework/service"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func ClearOperationRecord() error {
	var operationRecords []model.OperationRecord
	var ids []uint
	err := global.DB.Model(&model.OperationRecord{}).Order("created_at asc").Limit(3).Find(&operationRecords).Error
	if err != nil {
		return err
	}
	for _, record := range operationRecords {
		ids = append(ids, uint(record.ID))
	}
	return service.OperationRecordServiceApp.DeleteOperationRecordByIds(ids)
}

var cronStr = "*/30 * * * * *"
var clearOperationRecordID cron.EntryID

func AddClerOperationRecordTask(cron *cron.Cron) {
	var err error
	clearOperationRecordID, err = cron.AddFunc(cronStr, func() {
		fmt.Println("定时任务开始...")
		//err := ClearOperationRecord()
		ExecuteWithRetry(ClearOperationRecord, 3)
		fmt.Println("定时任务运行失败")

		fmt.Println("定时任务结束...")
	})
	if err != nil {
		panic(err)
	}
}

func ExecuteWithRetry(job func() error, maxRetries int) {
	for i := 0; i < maxRetries; i++ {
		err := job()
		if err == nil {
			return
		}
		fmt.Printf("第%d 次尝试失败，：%v", i+1, err.Error())
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("%d次尝试均失败", maxRetries)
	global.Cron.Remove(clearOperationRecordID)
}
