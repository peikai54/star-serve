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

	ProjectRoute(r)
}

func ProjectRoute(r *gin.Engine) {
	project := r.Group("/project")

	project.POST("/deleted", controller.DeleteProject)
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

	static := r.Group("/static")

	static.StaticFile("/resume", "static/.resume.pdf")

	static.StaticFile("/express", "static/express.zip")

	static.StaticFile("/vue-template", "static/vue-template.zip")

	static.StaticFile("/react-js", "static/react-js.zip")

	static.StaticFile("/react-ts", "static/react-ts.zip")

	static.StaticFile("/gin-rest", "static/gin-rest.zip")

	static.StaticFile("/koa2", "static/koa2.zip")
}
