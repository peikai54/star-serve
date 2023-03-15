package server

import (
	"serve/dto"
	"serve/model"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context, data dto.StoryAddReq) error {
	story := model.StroyModel{
		StoryName: data.StoryName,
		ProjectId: data.ProjectId,
		Status:    data.Status,
		CreatedBy: data.CreatedBy,
		// Handler:   data.Handler,
	}
	err := model.CreateStory(story)
	return err
}
