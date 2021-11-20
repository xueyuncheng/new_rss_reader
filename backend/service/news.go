package service

import (
	"context"
	"myNewFeed/cache"
	"myNewFeed/database"
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"net/http"
	"net/url"
)

func ListNews(ctx context.Context, req *model.ListNewsReq) ([]*model.News, error) {
	if req.FeedID != 0 {
		return database.ListNewsByFeedID(ctx, req.FeedID)
	}
	return cache.ListNews(ctx)
}

func RefreshNews() {
	ctx := context.Background()

	feeds, err := cache.ListFeed(ctx)
	if err != nil {
		return
	}

	proxyUrl, err := url.Parse("http://127.0.0.1:8123")
	if err != nil {
		log.Sugar.Errorw("parse proxy url error", "error", err)
		return
	}

	news := make([]*model.News, 0, 128)
	for _, v := range feeds {
		lastNewsTime, err := cache.GetLastNewsTime(ctx, v.ID)
		if err != nil {
			return
		}

		feed, err := srv.FeedParser.ParseURL(v.Name)
		if err != nil {
			log.Sugar.Errorf("ParseURL %v error: %v", v.Name, err)

			srv.FeedParser.Client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			}
			continue
		}

		for _, v2 := range feed.Items {
			if !v2.PublishedParsed.After(lastNewsTime) {
				continue
			}

			tmp := &model.News{
				Title:       v2.Title,
				Link:        v2.Link,
				PublishTime: *v2.PublishedParsed,
				FeedID:      int(v.ID),
				FeedName:    v.Name,
			}
			news = append(news, tmp)
		}

		if err := cache.SetLastNewsTime(ctx, v.ID); err != nil {
			return
		}
	}

	if err := database.AddNews(ctx, news...); err != nil {
		return
	}

}
