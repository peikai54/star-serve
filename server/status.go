package server

import (
	"serve/dto"
	"serve/model"

	"github.com/gin-gonic/gin"
)

func GetStatusList(c *gin.Context) ([]dto.StatusInfo, error) {
	statusList, err := model.GetStatusList()
	var result = []dto.StatusInfo{}
	for _, status := range statusList {
		result = append(result, dto.StatusInfo{StatusName: status.StatusName})
	}
	return result, err
}
