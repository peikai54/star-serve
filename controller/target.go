package controller

import (
	"fmt"
	"net/http"
	"serve/dto"
	"serve/server"

	"github.com/gin-gonic/gin"
)

func AddTarget(c *gin.Context) {
	var data dto.AddTargetReq
	var err error = nil
	err = c.Bind(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "入参错误"})
		return
	}

	err = server.AddTarget(c, data)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "插入成功"})
}

func GetTargetList(c *gin.Context) {
	list, err := server.GetTargetList(c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"list": list})
}
