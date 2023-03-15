package model

import (
	"serve/dto"
	"time"
)

type StroyModel struct {
	StoryId   int64 `gorm:"primaryKey"`
	ProjectId int64
	StoryName string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy int64
	Status    string
	// Handler   int64
}

func (StroyModel) TableName() string {
	return "story"
}

func StoryList(data dto.StoryListReq) {

}

// 创建需求
func CreateStory(story StroyModel) error {
	err := DbConnect.Create(&story).Error
	return err
}

// func StorytList(data dto.ProjectListReq) ([]ProjectModel, error) {
// 	projectList := []StroyModel{}

// 	query := DbConnect

// 	if data.Name != "" {
// 		query = query.Where("project_name LIKE BINARY ?", "%"+data.Name+"%")
// 	}

// 	if data.Type > 0 {
// 		query = query.Where("project_type = ?", data.Type)
// 	}

// 	if data.CreateAtEnd > 0 && data.CreateAtStart > 0 {
// 		create_at := time.Unix(data.CreateAtStart/1000, 0)
// 		end_at := time.Unix(data.CreateAtEnd/1000, 0)
// 		query = query.Where("create_at BETWEEN ? AND ?", create_at, end_at)
// 	}

// 	result := query.Find(&projectList)
// 	return projectList, result.Error
// }
