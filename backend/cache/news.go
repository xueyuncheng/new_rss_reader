package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"myNewFeed/database"
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"time"

	"github.com/go-redis/redis/v8"
)

func ListNews(ctx context.Context) ([]*model.News, error) {
	newsBytes, err := rdb.Get(ctx, "news").Bytes()
	if err == redis.Nil {
		news, err := database.ListNews(ctx)
		if err != nil {
			return nil, err
		}

		tmp, err := json.Marshal(news)
		if err != nil {
			log.Sugar.Errorw("cache.GetNews", "error", err)
			return nil, fmt.Errorf("cache.GetNews: %w", err)
		}

		if err := rdb.Set(ctx, "news", tmp, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("cache.GetNews", "error", err)
			return nil, fmt.Errorf("cache.GetNews: %w", err)
		}

		return ListNews(ctx)
	} else if err != nil {
		log.Sugar.Errorw("cache.GetNews", "error", err)
		return nil, fmt.Errorf("cache.GetNews: %w", err)
	}

	news := make([]*model.News, 0, 1024)
	if err := json.Unmarshal(newsBytes, &news); err != nil {
		log.Sugar.Errorw("cache.GetNews", "error", err)
		return nil, fmt.Errorf("cache.GetNews: %w", err)
	}

	return news, nil
}

func ListNewsByFeedID(ctx context.Context, req *model.ListNewsReq) ([]*model.News, error) {
	newsBytes, err := rdb.Get(ctx, fmt.Sprintf("feed:%v", req.FeedID)).Bytes()
	if err == redis.Nil {
		news, err := database.ListNewsByFeedID(ctx, req.FeedID)
		if err != nil {
			log.Sugar.Errorw("cache.ListNewsByFeedID", "error", err)
			return nil, fmt.Errorf("cache.ListNewsByFeedID: %w", err)
		}

		newsBytes, err = json.Marshal(news)
		if err != nil {
			log.Sugar.Errorw("cache.ListNewsByFeedID", "error", err)
			return nil, fmt.Errorf("cache.ListNewsByFeedID: %w", err)
		}

		key := fmt.Sprintf("feed:%v", req.FeedID)
		if err := rdb.Set(ctx, key, newsBytes, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("cache.ListNewsByFeedID", "error", err)
			return nil, fmt.Errorf("cache.ListNewsByFeedID: %w", err)
		}

		return ListNewsByFeedID(ctx, req)
	} else if err != nil {
		log.Sugar.Errorw("cache.ListNewsByFeedID", "error", err)
		return nil, fmt.Errorf("cache.ListNewsByFeedID: %w", err)
	}

	news := make([]*model.News, 0, 1024)
	if err := json.Unmarshal(newsBytes, &news); err != nil {
		log.Sugar.Errorw("cache.ListNewsByFeedID", "error", err)
		return nil, fmt.Errorf("cache.ListNewsByFeedID: %w", err)
	}

	return news, nil
}

func GetLastNewsTime(ctx context.Context, feedID int) (time.Time, error) {
	key := fmt.Sprintf("last.news.time:%v", feedID)
	lastNewsTimeString, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		lastNewsTime, err := database.GetLastNewsTime(ctx, feedID)
		if err != nil {
			return time.Time{}, err
		}

		value := lastNewsTime.Format("2006-01-02 15:04:05")
		if err := rdb.Set(ctx, key, value, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
			return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
		}

		return GetLastNewsTime(ctx, feedID)
	} else if err != nil {
		log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
		return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
	}

	lastNewsTime, err := time.ParseInLocation("2006-01-02 15:04:05", lastNewsTimeString, time.Local)
	if err != nil {
		log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
		return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
	}

	return lastNewsTime, nil
}

func SetLastNewsTime(ctx context.Context, feedID int) error {
	key := fmt.Sprintf("last.news.time:%v", feedID)
	value := time.Now().Format("2006-01-02 15:04:05")
	if err := rdb.Set(ctx, key, value, time.Hour).Err(); err != nil {
		log.Sugar.Errorw("cache.SetLastNewsTime", "error", err)
		return fmt.Errorf("cache.SetLastNewsTime: %w", err)
	}

	return nil
}

func DeleteNews(ctx context.Context) error {
	pipe := rdb.TxPipeline()
	defer pipe.Close()

	pipe.Del(ctx, "news")

	feeds, err := database.ListFeed(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		key := fmt.Sprintf("feed:%v", feed.ID)
		pipe.Del(ctx, key)
	}

	if _, err := pipe.Exec(ctx); err != nil {
		log.Sugar.Errorw("pipeline执行错误", "error", err)
		return fmt.Errorf("pipeline执行错误: %w", err)
	}

	return nil
}
