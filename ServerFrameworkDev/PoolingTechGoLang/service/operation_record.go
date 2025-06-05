package service

import (
	"ServerFramework/global"
	"ServerFramework/model"
)

type OperationRecordService struct{}

var OperationRecordServiceApp = new(OperationRecordService)

func (operationRecordService *OperationRecordService) CreateOperationRecord(OperationRecord model.OperationRecord) (err error) {
	err = global.DB.Create(&OperationRecord).Error
	return err
}

// 物理删除
func (operationRecordService *OperationRecordService) DeleteOperationRecordByIds(ids []uint) (err error) {
	err = global.DB.Unscoped().Delete(&[]model.OperationRecord{}, "id in (?)", ids).Error
	return err
}

func (operationRecordService *OperationRecordService) DeleteOperationRecord(OperationRecord model.OperationRecord) (err error) {
	err = global.DB.Delete(&OperationRecord).Error
	return err
}

func (operationRecordService *OperationRecordService) GetOperationRecord(id uint) (OperationRecord model.OperationRecord, err error) {
	err = global.DB.Where("id = ?", id).First(&OperationRecord).Error
	return
}
