package controller

import (
	"fmt"
	"net/http"
	"serve/dto"
	"serve/server"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var data dto.AddProjectReq
	var err error

	err = c.Bind(&data)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "入参错误"})
		return
	}

	err = server.CreateProject(c, data)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "创建成功"})
}

func GetProjectList(c *gin.Context) {
	var data dto.ProjectListReq
	var list []dto.ProjectListResp
	var err error

	err = c.ShouldBindQuery(&data)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "入参错误"})
		return
	}
	list, err = server.GetProjectList(c, data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"list": list})
}

func DeleteProject(c *gin.Context) {
	var data dto.DeletedProject
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "入参错误"})
		return
	}
	err = server.DeletedProject(c, data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

func UpdateProject(c *gin.Context) {
	var data dto.UpdateProjectReq
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "入参错误"})
		return
	}
	err = server.UpdateProject(c, data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "编辑成功"})
}
