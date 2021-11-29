package model

type AddFeedReq struct {
	Name string `json:"name" binding:"required"`
}

type ListFeedResp struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ErrorMsg string `json:"error_msg"`
}

type Feed struct {
	Model
	Name     string `json:"name"`
	ErrorMsg string `json:"error_msg"`
}

func (r *Feed) TableName() string {
	return "feed"
}
