package controller

import (
	"fmt"
	"serve/dto"
	"serve/server"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data *dto.LoginReq
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "入参错误"})
		return
	}
	token, err2 := server.Login(c, data)
	if err2 != nil {
		c.JSON(500, gin.H{"message": err2.Error()})
		return
	}
	c.SetCookie("token", token, 3600, "/", "", false, false)
	c.JSON(200, gin.H{"message": "登录成功", "token": token})
}

func GetUserInfo(c *gin.Context) {

	var data dto.UserInfoReq

	err := c.ShouldBindQuery(&data)

	if err != nil {
		c.JSON(403, gin.H{"message": "暂无权限"})
		return
	}

	userInfo, err2 := server.GetUserInfo(c, data.Token)
	if err2 != nil {
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	}
	c.JSON(200, gin.H{"info": userInfo})
}

func GetUserList(c *gin.Context) {
	result, err := server.GetUserList(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "获取用户列表错误")
		return
	}
	c.JSON(200, gin.H{"list": result})
}
