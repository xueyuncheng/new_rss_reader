package service

import (
	"myNewFeed/model"

	"github.com/mmcdole/gofeed"
)

var srv *model.Service

func InitService() {
	srv = &model.Service{
		FeedParser: gofeed.NewParser(),
	}
}
