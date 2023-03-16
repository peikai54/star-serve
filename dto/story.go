package dto

type StoryListReq struct {
	StoryName string `json:"story_name"`
	Handler   string `json:"handler"`
	CreatedBy string `json:"created_by"`
	Status    string `json:"status"`
	ProjectId int64  `json:"project_id"`
}

type StoryAddReq struct {
	StoryName string  `json:"story_name" binding:"required"`
	Handler   []int64 `json:"handler"`
	CreatedBy int64   `json:"created_by" binding:"required"`
	Status    string  `json:"status"`
	ProjectId int64   `json:"project_id" binding:"required"`
}
