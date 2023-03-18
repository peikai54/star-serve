package dto

type AddProjectReq struct {
	ProjectName string `json:"project_name" binding:"required"`
	ProjectType int64  `json:"project_type" binding:"required"`
	UserId      int64  `json:"user_id" binding:"required"`
}

type ProjectListReq struct {
	Ids           []int64 `form:"id"`
	Name          string  `form:"name"`
	Type          int64   `form:"type"`
	Creator       string  `form:"creator"`
	CreateAtStart int64   `form:"start_at"`
	CreateAtEnd   int64   `form:"end_at"`
}

type ProjectListResp struct {
	ProjectName string `json:"project_name"`
	ProjectType int64  `json:"project_type"`
	ProjectId   int64  `json:"project_id"`
	CreatOrId   int64  `json:"creator_id"`
	CreatOrName string `json:"creator_name"`
	CreateAt    string `json:"create_at"`
}

type DeletedProject struct {
	ProjectId int64 `json:"project_id"`
}
