package service

import (
	"context"
	"myNewFeed/cache"
	"myNewFeed/database"
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"time"

	"github.com/mmcdole/gofeed"
)

func ListNews(ctx context.Context, req *model.ListNewsReq) ([]*model.News, error) {
	if req.FeedID != 0 {
		return cache.ListNewsByFeedID(ctx, req)
	}

	return cache.ListNews(ctx)
}

func RefreshNews() {
	srv.Mutex.Lock()
	defer srv.Mutex.Unlock()

	ctx := context.Background()
	feeds, err := cache.ListFeed(ctx)
	if err != nil {
		return
	}

	newses := make([]*model.News, 0, 128)
	for _, v := range feeds {
		news, err := getFeedNews(ctx, v)
		if err != nil {
			v.ErrorMsg = err.Error()
			if err := database.UpdateFeed(ctx, v); err != nil {
				return
			}
		}

		if err := updateFeedStatus(ctx, v.ID); err != nil {
			return
		}

		newses = append(newses, news...)
	}

	if err := database.AddNews(ctx, newses...); err != nil {
		return
	}
}

func getFeedNews(ctx context.Context, feedM *model.Feed) ([]*model.News, error) {
	lastNewsTime, err := cache.GetLastNewsTime(ctx, feedM.ID)
	if err != nil {
		return nil, err
	}

	feed, err := srv.FeedParser.ParseURL(feedM.Name)
	if err != nil {
		log.Sugar.Errorf("ParseURL %v error: %v", feedM.Name, err)
		return nil, err
	}

	log.Sugar.Infow("ParseURL succeed", "url", feedM.Name)

	newses := make([]*model.News, 0, 1024)
	for _, v := range feed.Items {
		news, ok := getNews(feed, feedM, v, lastNewsTime)
		if !ok {
			continue
		}

		newses = append(newses, news)
	}

	if len(newses) > 0 {
		if err := cache.DeleteFeedByID(ctx, feedM.ID); err != nil {
			return nil, err
		}
	}

	if err := cache.SetLastNewsTime(ctx, feedM.ID); err != nil {
		return nil, err
	}

	return newses, nil
}

func getNews(feed *gofeed.Feed, feedM *model.Feed, item *gofeed.Item, lastNewsTime time.Time) (*model.News, bool) {
	if item.PublishedParsed != nil && !item.PublishedParsed.After(lastNewsTime) {
		return nil, false
	}

	if feed.UpdatedParsed != nil && !feed.UpdatedParsed.After(lastNewsTime) {
		return nil, false
	}

	var publishTime time.Time
	if item.PublishedParsed != nil {
		publishTime = *item.PublishedParsed
	} else if feed.UpdatedParsed != nil {
		publishTime = *feed.UpdatedParsed
	} else {
		log.Sugar.Errorw("feed has no update time and news has no publish time", "feed_name", feedM.Name)
		return nil, false
	}

	news := &model.News{
		Title:       item.Title,
		Link:        item.Link,
		PublishTime: publishTime,
		FeedID:      feedM.ID,
		FeedName:    feedM.Name,
	}

	return news, true
}

func DeleteOldNews() {
	srv.Mutex.Lock()
	defer srv.Mutex.Unlock()

	ctx := context.Background()
	if err := database.DeleteOldNews(ctx); err != nil {
		return
	}

	if err := cache.DeleteNews(ctx); err != nil {
		return
	}

	if err := cache.DeleteFeedNews(ctx); err != nil {
		return
	}
}

func StatNews(ctx context.Context) (*model.Chart, error) {
	news, err := cache.ListNews(ctx)
	if err != nil {
		return nil, err
	}

	m := make(map[string]int)
	for _, v := range news {
		m[v.PublishTime.Format("01-02")]++
	}

	chart := &model.Chart{
		Name:  "news",
		Items: []*model.Item{},
	}

	if len(news) == 0 {
		return chart, nil
	}

	newest := news[0].PublishTime.Truncate(24 * time.Hour)
	oldest := news[len(news)-1].PublishTime.Truncate(24 * time.Hour)
	for day := oldest; !day.After(newest); day = day.AddDate(0, 0, 1) {
		tmp := &model.Item{
			Name:  day.Format("01-02"),
			Value: float64(m[day.Format("01-02")]),
		}

		chart.Items = append(chart.Items, tmp)
	}

	return chart, nil
}

func updateFeedStatus(ctx context.Context, id int) error {
	ok, err := cache.GetFeedStatus(ctx, id)
	if err != nil {
		return err
	}

	if ok {
		if err := cache.DeleteFeed(ctx); err != nil {
			return err
		}

		if err := cache.SetFeedStatus(ctx, id, false); err != nil {
			return err
		}
	} else {
		if err := database.UpdateFeedStatus(ctx, id, ""); err != nil {
			return err
		}

		if err := cache.DeleteFeed(ctx); err != nil {
			return err
		}

		if err := cache.SetFeedStatus(ctx, id, true); err != nil {
			return err
		}
	}

	return nil
}
