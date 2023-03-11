package controller

import (
	"fmt"
	"serve/dto"
	"serve/server"

	"github.com/gin-gonic/gin"
)

func AddCheckRecord(c *gin.Context) {
	var data dto.AddCheckRecordReq
	var err error
	err = c.Bind(&data)
	if err != nil {
		c.JSON(400, gin.H{"messsage": "入参错误"})
		return
	}
	err = server.AddCheckRecord(c, data)
	if err != nil {
		c.JSON(500, gin.H{"messsage": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "插入成功"})
}

func DeleteRecord(c *gin.Context) {
	var data dto.DeletedCheckRecordReq
	var err error
	err = c.Bind(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": "入参错误"})
		return
	}
	err = server.DeletedCheckRecord(c, data)
	if err != nil {
		c.JSON(500, gin.H{"messsage": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}
