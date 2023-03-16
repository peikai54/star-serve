package server

import (
	"serve/dto"
	"serve/model"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context, data dto.StoryAddReq) error {
	story := model.StoryModel{
		StoryName: data.StoryName,
		ProjectId: data.ProjectId,
		Status:    data.Status,
		CreatedBy: data.CreatedBy,
	}
	// 正常插入
	if data.Handler == nil {
		err := model.CreateStory(story)
		return err
	}

	err := model.CreateStoryWithHandler(story, data.Handler)

	return err
}
