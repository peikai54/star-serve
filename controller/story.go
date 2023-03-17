package controller

import (
	"fmt"
	"serve/dto"
	"serve/server"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	var data dto.StoryAddReq
	var err error
	err = c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	err = server.CreateStory(c, data)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "新建成功"})
}

func StoryList(c *gin.Context) {
	var data dto.StoryListReq
	var err error
	err = c.ShouldBindQuery(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	var storyList []dto.StoryListResp
	storyList, err = server.StoryList(c, data)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"list": storyList})
}

func GetStatusList(c *gin.Context) {
	result, err := server.GetStatusList(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "需求状态获取错误"})
		return
	}
	c.JSON(200, gin.H{"list": result})
}
