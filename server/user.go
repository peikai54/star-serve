package server

import (
	"serve/dto"
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
