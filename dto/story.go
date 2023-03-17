package dto

type StoryListReq struct {
	StoryName string `form:"story_name"`
	Handler   int64  `form:"handler"`
	CreatedBy string `form:"created_by"`
	Status    string `form:"status"`
	ProjectId int64  `form:"project_id"`
}

type StoryAddReq struct {
	StoryName string  `json:"story_name" binding:"required"`
	Handler   []int64 `json:"handler"`
	CreatedBy int64   `json:"created_by" binding:"required"`
	Status    string  `json:"status"`
	ProjectId int64   `json:"project_id" binding:"required"`
}

type StoryListResp struct {
	StoryId     int64       `json:"story_id"`
	StoryName   string      `json:"story_name"`
	Handler     []UserInfo  `json:"handler_list"`
	CreatedBy   UserInfo    `json:"created_by"`
	Status      string      `json:"status"`
	ProjectInfo ProjectInfo `json:"project_info"`
}

type UserInfo struct {
	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
}

type ProjectInfo struct {
	ProjectName string `json:"project_name"`
	ProjectId   int64  `json:"project_id"`
}

type StoryHandlerInfo struct {
	Name    string `json:"name"`
	UserId  int64  `json:"user_id"`
	StoryId int64  `json:"story_id"`
}

type StatusInfo struct {
	StatusName string `json:"status_name"`
}
