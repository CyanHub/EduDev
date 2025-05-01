package task

import (
	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/model"
	"github.com/CyanHub/EduDev/service"
)

var operationRecordService = service.OperationRecordServiceApp

func ClearOperationRecord() {
	var ids []uint
	var records []model.OperationRecord
	// 按创建时间升序排序
	global.DB.Model(&model.OperationRecord{}).Order("created_at asc").Limit(10).Find(&records)
	for _, record := range records {
		ids = append(ids, uint(record.ID))
	}
	operationRecordService.DeleteOperationRecordByIds(ids)
}
