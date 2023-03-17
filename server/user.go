package server

import (
	"serve/dto"
	"serve/model"
	"serve/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, data *dto.LoginReq) (string, error) {
	var err error
	var token string
	if data.UserName == "游客测试账号" {
		token, err = utils.GenerateToken(data.UserName, "")
		return token, err
	}
	return token, err
}

func GetUserInfo(c *gin.Context, token string) (dto.UserInfoResp, error) {
	config, err := utils.ParseToken(token)
	var userInfo = dto.UserInfoResp{}
	if err != nil {
		return userInfo, err
	}
	if config.Username == "游客测试账号" {
		userInfo.UserName = config.Username
		userInfo.UserId = 2
	}
	return userInfo, err
}

func GetUserList(c *gin.Context) ([]dto.UserInfoResp, error) {
	var result = []dto.UserInfoResp{}
	userList, err := model.GetUserList()
	for _, user := range userList {
		item := dto.UserInfoResp{UserName: user.UserName, UserId: user.UserId}
		result = append(result, item)
	}
	return result, err
}
