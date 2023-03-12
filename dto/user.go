package dto

type LoginReq struct {
	Password string `json:"password" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
}
