package server

import (
	"serve/dto"
	"serve/model"
	"serve/utils"

	"github.com/gin-gonic/gin"
)

func AddCheckRecord(c *gin.Context, data dto.AddCheckRecordReq) error {
	theEnd := utils.GetTodayEnd()
	var record = model.CheckRecordModel{
		CheckType: data.CheckType,
		TargetId:  data.Id,
		CheckAt:   theEnd,
	}
	err := model.CreateCheckRecord(record)
	return err
}

func DeletedCheckRecord(c *gin.Context, data dto.DeletedCheckRecordReq) error {
	var record = model.CheckRecordModel{
		Id: data.Id,
	}
	err := model.DeleteCheckRecord(record)
	return err
}
