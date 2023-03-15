package router

import (
	"net/http"
	"serve/controller"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func ResignRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.POST("/save-target", controller.AddTarget)

	r.GET("/target-list", controller.GetTargetList)

	r.POST("/add-check-record", controller.AddCheckRecord)

	r.DELETE("/delete-check-record", controller.DeleteRecord)

	r.POST("/add-project", controller.CreateProject)

	r.GET("/projects", controller.GetProjectList)

	r.POST("/login", controller.Login)

	r.GET("/user-info", controller.GetUserInfo)

	r.POST("/add-story", controller.CreateStory)
}
