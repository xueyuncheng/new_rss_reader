package model

import (
	"net/url"

	"github.com/mmcdole/gofeed"
)

type Service struct {
	FeedParser *gofeed.Parser
	ProxyUrl   *url.URL
}
