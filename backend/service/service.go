package service

import (
	"fmt"
	"myNewFeed/model"
	"net/url"

	"github.com/mmcdole/gofeed"
)

var srv *model.Service

func InitService(config *model.Proxy) {
	srv = &model.Service{
		FeedParser: gofeed.NewParser(),
		ProxyUrl: &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%v:%v", config.Host, config.Port),
		},
	}
}
