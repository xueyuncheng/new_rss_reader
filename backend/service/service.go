package service

import (
	"myNewFeed/model"
	"sync"

	"github.com/mmcdole/gofeed"
)

var srv *model.Service

func InitService() {
	srv = &model.Service{
		FeedParser: gofeed.NewParser(),
		Mutex:      &sync.Mutex{},
	}
}
