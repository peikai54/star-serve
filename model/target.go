package model

import (
	"fmt"
	"time"
)

type TargetModel struct {
	ID         int64 `gorm:"primaryKey"`
	Name       string
	Desc       string
	Type       int64
	StartAt    time.Time
	EndAt      time.Time
	Status     int64
	CheckCount int64
	FailCount  int64
}

func (TargetModel) TableName() string {
	return "target"
}

// 创建target
func CreateTarget(target TargetModel) error {
	err := DbConnect.Create(&target).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func TargetList() ([]TargetModel, error) {
	targets := []TargetModel{}
	result := DbConnect.Find(&targets)
	return targets, result.Error
}
