package model

import "time"

type ListNewsReq struct {
	FeedID int `form:"feed_id"`
}

type ListNewsResp struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	PublishTime string `json:"publish_time"`
	FeedID      int    `json:"feed_id"`
	FeedName    string `json:"feed_name"`
}

type News struct {
	Model
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	PublishTime time.Time `json:"publish_time"`
	FeedID      int       `json:"feed_id"`
	FeedName    string    `json:"feed_name"`
}

func (f *News) TableName() string {
	return "news"
}
