package model

import "serve/dto"

type StroyHandlerModel struct {
	StoryId int64
	Handler int64
}

func BatchGetHandlerByStoryIds(storyIds []int64) ([]dto.StoryHandlerInfo, error) {
	var handlerList = []StroyHandlerModel{}
	var userList = []UserModel{}
	var err error
	err = DbConnect.Where("story_id in ?", storyIds).Find(&handlerList).Error

	if err != nil {
		return nil, err
	}

	var userIds []int64
	for _, item := range handlerList {
		userIds = append(userIds, item.Handler)
	}
	userList, err = BatchGetUserByIds(userIds, "")
	if err != nil {
		return nil, err
	}

	var result = []dto.StoryHandlerInfo{}

	for _, item := range handlerList {
		for _, user := range userList {
			if user.UserId == item.Handler {
				item := dto.StoryHandlerInfo{Name: user.UserName, UserId: item.Handler, StoryId: item.StoryId}
				result = append(result, item)
			}
		}
	}

	return result, err
}

func CreateStoryHandlerInBatch(list []StroyHandlerModel) error {
	err := DbConnect.CreateInBatches(list, 100).Error
	return err
}

func (StroyHandlerModel) TableName() string {
	return "story_connect"
}
