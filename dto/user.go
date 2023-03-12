package dto

type LoginReq struct {
	Password string `json:"password"`
	UserName string `json:"user_name" binding:"required"`
}

type UserInfoResp struct {
	UserName string `json:"user_name"`
	UserId   int64  `json:"user_id"`
}
