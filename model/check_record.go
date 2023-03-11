package model

import "time"

type CheckRecordModel struct {
	TargetId  int64
	CheckType int64
	CheckAt   time.Time
	Id        int64 `gorm:"primaryKey"`
}

func (CheckRecordModel) TableName() string {
	return "check_record"
}

func CreateCheckRecord(record CheckRecordModel) error {
	err := DbConnect.Create(&record).Error
	return err
}

func DeleteCheckRecord(record CheckRecordModel) error {
	err := DbConnect.Delete(&record).Error
	return err
}
