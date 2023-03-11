package server

import (
	"serve/dto"
	"serve/model"
	"serve/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTarget(c *gin.Context, data dto.AddTargetReq) error {
	var targetModel = model.TargetModel{
		Name:       data.Name,
		Type:       data.Type,
		Desc:       data.Desc,
		Status:     utils.TargetRunning,
		CheckCount: 0,
		FailCount:  0,
		StartAt:    time.Unix(data.StartAt, 0),
		EndAt:      time.Unix(data.EndAt, 0),
	}
	err := model.CreateTarget(targetModel)
	return err
}

func GetTargetList(c *gin.Context) ([]dto.TargetListResp, error) {
	targets, err := model.TargetList()
	var list = []dto.TargetListResp{}

	for _, target := range targets {
		var item = dto.TargetListResp{
			Name:       target.Name,
			Type:       target.Type,
			StartAt:    target.StartAt.Format("2006-01-01"),
			Desc:       target.Desc,
			EndAt:      target.EndAt.Format("2006-01-01"),
			FailCount:  target.FailCount,
			CheckCount: target.CheckCount,
			Status:     target.Status,
		}
		list = append(list, item)
	}
	return list, err
}
