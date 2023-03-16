package model

import (
	"serve/dto"
	"time"
)

type StoryModel struct {
	StoryId   int64 `gorm:"primaryKey"`
	ProjectId int64
	StoryName string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy int64
	Status    string
	// Handler   int64
}

func (StoryModel) TableName() string {
	return "story"
}

func StoryList(data dto.StoryListReq) {

}

// 创建需求
func CreateStory(story StoryModel) error {
	err := DbConnect.Create(&story).Error
	return err
}

func CreateStoryWithHandler(story StoryModel, handlerList []int64) error {
	// 创建事务
	tx := DbConnect.Begin()
	var err error = nil

	err = tx.Create(&story).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	var storyHandler = []StroyHandlerModel{}

	for _, id := range handlerList {
		item := StroyHandlerModel{Handler: id, StoryId: story.StoryId}
		storyHandler = append(storyHandler, item)
	}

	err = tx.CreateInBatches(storyHandler, 100).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err

}
