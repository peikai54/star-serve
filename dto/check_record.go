package dto

type AddCheckRecordReq struct {
	Id        int64 `json:"target_id" binding:"required"`
	CheckType int64 `json:"check_type" binding:"required"`
}

type DeletedCheckRecordReq struct {
	Id int64 `json:"id" binding:"required"`
}
