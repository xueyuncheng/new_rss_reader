package model

import "time"

type AddFeedReq struct {
	Name string `json:"name" binding:"required"`
}

type Feed struct {
	Model
	Name string `json:"name"`
}

func (r *Feed) TableName() string {
	return "feed"
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
