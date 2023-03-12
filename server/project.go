package server

import (
	"serve/dto"
	"serve/model"

	"github.com/gin-gonic/gin"
)

// 创建项目
func CreateProject(c *gin.Context, data dto.AddProjectReq) error {
	var projcetModel = model.ProjectModel{
		ProjectName: data.ProjectName,
		ProjectType: data.ProjectType,
		Creator:     data.UserId,
	}
	err := model.CreateProject(projcetModel)
	return err
}

func GetProjectList(c *gin.Context, data dto.ProjectListReq) ([]dto.ProjectListResp, error) {
	targets, err := model.ProjectList(data)

	if err != nil {
		return nil, err
	}
	var list = []dto.ProjectListResp{}
	var ids = []int64{}
	for _, target := range targets {
		ids = append(ids, target.Creator)
	}

	users, err2 := model.BatchGetUserByIds(ids, data.Creator)

	if err2 != nil {
		return nil, err2
	}

	for _, target := range targets {
		item := dto.ProjectListResp{
			ProjectName: target.ProjectName,
			ProjectType: target.ProjectType,
			ProjectId:   target.ProjectId,
			CreatOrId:   target.Creator,
			CreateAt:    target.CreateAt.Format("2006-01-02 15:04"),
		}
		for _, user := range users {
			if user.UserId == item.CreatOrId {
				item.CreatOrName = user.UserName
				list = append(list, item)
			}
		}
	}

	return list, err
}
