package model

type AddFeedReq struct {
	Name string `json:"name" binding:"required"`
}

type GetFeedResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Feed struct {
	Model
	Name string `json:"name"`
}

func (r *Feed) TableName() string {
	return "feed"
}
