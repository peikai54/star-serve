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

	StoryRoute(r)

	UserRoute(r)

	StaticConfig(r)
}

func UserRoute(r *gin.Engine) {
	user := r.Group("/user")

	user.GET("/list", controller.GetUserList)
}

func StoryRoute(r *gin.Engine) {
	story := r.Group("/story")

	story.POST("/add", controller.CreateStory)
	story.GET("/list", controller.StoryList)
	story.GET("/status", controller.GetStatusList)
}

func StaticConfig(r *gin.Engine) {

	// static := r.Group("/static")

	// r.Static("/static", "/static")
	r.StaticFile("/static/resume", "static/resume.pdf")
}
