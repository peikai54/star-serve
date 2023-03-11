package model

import (
	"fmt"
	"time"
)

type ProjectModel struct {
	ProjectId   int64 `gorm:"primaryKey"`
	ProjectName string
	Creator     int64
	CreateAt    time.Time `gorm:"autoCreateTime"`
	ProjectType int64
}

func (ProjectModel) TableName() string {
	return "project"
}

// 创建项目
func CreateProject(project ProjectModel) error {
	err := DbConnect.Create(&project).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func ProjectList() ([]ProjectModel, error) {
	projectList := []ProjectModel{}
	result := DbConnect.Find(&projectList)
	return projectList, result.Error
}
