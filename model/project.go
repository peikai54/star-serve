package model

import (
	"fmt"
	"serve/dto"
	"time"

	"gorm.io/gorm"
)

type ProjectModel struct {
	ProjectId   int64 `gorm:"primaryKey"`
	ProjectName string
	Creator     int64
	CreateAt    time.Time `gorm:"autoCreateTime"`
	ProjectType int64
	IsDeleted   int64
}

func (ProjectModel) TableName() string {
	return "project"
}

// 删除项目
func DeleteProject(tx *gorm.DB, ProjectId int64) error {
	return tx.Model(&ProjectModel{}).Where("project_id = ?", ProjectId).Update("is_deleted", 1).Error
}

// 创建项目
func CreateProject(project ProjectModel) error {
	err := DbConnect.Create(&project).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func ProjectList(data dto.ProjectListReq) ([]ProjectModel, error) {
	projectList := []ProjectModel{}

	query := DbConnect

	if data.Name != "" {
		query = query.Where("project_name LIKE BINARY ?", "%"+data.Name+"%")
	}

	if data.Type > 0 {
		query = query.Where("project_type = ?", data.Type)
	}

	if data.CreateAtEnd > 0 && data.CreateAtStart > 0 {
		create_at := time.Unix(data.CreateAtStart/1000, 0)
		end_at := time.Unix(data.CreateAtEnd/1000, 0)
		query = query.Where("create_at BETWEEN ? AND ?", create_at, end_at)
	}

	if data.Ids != nil {
		query = query.Where("project_id in ?", data.Ids)
	}

	query = query.Where("is_deleted = ?", 0)

	result := query.Find(&projectList)
	return projectList, result.Error
}

type ProjectModelRepo interface {
	UpdateProject(tx *gorm.DB, data dto.UpdateProjectReq) error
}

type ProjectModelStruct struct{}

func (ProjectModelStruct) UpdateProject(tx *gorm.DB, data dto.UpdateProjectReq) error {
	return tx.Model(&ProjectModel{}).Where("project_id = ?", data.ProjectId).
		Updates(ProjectModel{ProjectName: data.ProjectName, ProjectType: data.ProjectType}).Error
}

var ProjectModelSet = ProjectModelStruct{}
