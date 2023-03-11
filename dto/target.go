package dto

type AddTargetReq struct {
	Name    string `json:"name" binding:"required"`
	Desc    string `json:"desc"`
	Type    int64  `json:"type" binding:"required"`
	StartAt int64  `json:"start_at" binding:"required"`
	EndAt   int64  `json:"end_at" binding:"required"`
}

type TargetListResp struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Type       int64  `json:"type"`
	StartAt    string `json:"start_at"`
	EndAt      string `json:"end_at"`
	Status     int64  `json:"status"`
	CheckCount int64  `json:"check_out"`
	FailCount  int64  `json:"fail_count"`
}
