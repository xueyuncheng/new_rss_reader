package model

import (
	"sync"

	"github.com/mmcdole/gofeed"
)

type Service struct {
	FeedParser *gofeed.Parser
	Mutex      *sync.Mutex
}
