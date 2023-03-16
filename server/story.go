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

func StoryList(c *gin.Context, data dto.StoryListReq) ([]dto.StoryListResp, error) {
	var tempList = []model.StoryModel{}
	var tempProjectList = []model.ProjectModel{}
	var tempCreatorList = []model.UserModel{}
	var tempHandlerList = []dto.StoryHandlerInfo{}
	var storyList []dto.StoryListResp
	var err error

	tempList, err = model.StoryList(data)

	if err != nil {
		return nil, err
	}

	// 获取项目列表内容
	var projectIds []int64
	var createdIds []int64
	var storyIds []int64

	for _, item := range tempList {
		projectIds = append(projectIds, item.ProjectId)
		createdIds = append(createdIds, item.CreatedBy)
		storyIds = append(storyIds, item.StoryId)
	}

	// 获取项目列表
	tempProjectList, err = model.ProjectList(dto.ProjectListReq{Ids: projectIds})

	if err != nil {
		return nil, err
	}

	// 获取用户列表
	tempCreatorList, err = model.BatchGetUserByIds(createdIds, data.CreatedBy)

	if err != nil {
		return nil, err
	}

	tempHandlerList, err = model.BatchGetHandlerByStoryIds(storyIds)

	for _, item := range tempList {

		var isAppend bool = false

		// 数据整合，添加项目信息
		var projectInfo = dto.ProjectInfo{ProjectId: item.ProjectId}
		for _, project := range tempProjectList {
			if project.ProjectId == item.ProjectId {
				projectInfo.ProjectName = project.ProjectName
			}
		}

		// 数据整合，添加创建人信息
		var creatorData = dto.UserInfo{UserId: item.CreatedBy}
		for _, creator := range tempCreatorList {
			if creator.UserId == item.CreatedBy {
				isAppend = true
				creatorData.Name = creator.UserName
			}
		}

		// 数据整合，添加处理人信息
		var handlerList = []dto.UserInfo{}
		for _, handler := range tempHandlerList {
			if handler.StoryId == item.StoryId {
				handlerList = append(handlerList, dto.UserInfo{Name: handler.Name, UserId: handler.UserId})
			}
		}

		story := dto.StoryListResp{
			StoryName:   item.StoryName,
			Status:      item.Status,
			ProjectInfo: projectInfo,
			StoryId:     item.StoryId,
			CreatedBy:   creatorData,
			Handler:     handlerList,
		}

		if isAppend {
			storyList = append(storyList, story)
		}
	}

	// 根据处理人进行过滤
	if data.Handler > 0 {
		for index, story := range storyList {
			var couldFind bool = false
			for _, item := range story.Handler {
				if item.UserId == data.Handler {
					couldFind = true
				}
			}
			if !couldFind {
				if len(storyList) == 1 {
					storyList = []dto.StoryListResp{}
				} else {
					storyList = storyList[:index+copy(storyList[index:], storyList[index+1:])]
				}
			}
		}
	}

	return storyList, err

}
