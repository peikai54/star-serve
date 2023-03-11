package dto

type AddProjectReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	ProjectType int64  `json:"project_type" binding:"required"`
	UserId      int64  `json:"user_id" binding:"required"`
}

type ProjectListResp struct {
	ProjectName string `json:"project_name"`
	ProjectType int64  `json:"project_type"`
	ProjectId   int64  `json:"project_id"`
	CreatOrId   int64  `json:"creator_id"`
	CreatOrName string `json:"creator_name"`
	CreateAt    string `json:"create_at"`
}
