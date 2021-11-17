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

func GetNews(ctx context.Context) ([]*model.News, error) {
	news := make([]*model.News, 0, 1024)
	var err error

	newsByte, err := RedisClient.Get(ctx, "news").Bytes()
	if err == redis.Nil {
		news, err = database.GetNews(ctx)
		if err != nil {
			return nil, err
		}

		tmp, err := json.Marshal(news)
		if err != nil {
			log.Sugar.Errorw("cache.GetNews", "error", err)
			return nil, fmt.Errorf("cache.GetNews: %w", err)
		}

		if err := RedisClient.Set(ctx, "news", tmp, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("cache.GetNews", "error", err)
			return nil, fmt.Errorf("cache.GetNews: %w", err)
		}

		return GetNews(ctx)
	} else if err != nil {
		log.Sugar.Errorw("cache.GetNews", "error", err)
		return nil, fmt.Errorf("cache.GetNews: %w", err)
	}

	if err := json.Unmarshal(newsByte, &news); err != nil {
		log.Sugar.Errorw("cache.GetNews", "error", err)
		return nil, fmt.Errorf("cache.GetNews: %w", err)
	}

	return news, nil
}

func AddNews(ctx context.Context, news *model.News) error {
	if err := database.DB.Create(news).Error; err != nil {
		log.Sugar.Errorw("cache.AddNews", "error", err)
		return fmt.Errorf("cache.AddNews: %w", err)
	}

	return nil
}

func GetLastNewsTime(ctx context.Context, feedID int) (time.Time, error) {
	var lastNewsTime time.Time
	var err error

	key := fmt.Sprintf("lastNewsTime:%v", feedID)

	lastNewsTimeByte, err := RedisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		lastNewsTime, err = database.GetLastNewsTime(ctx, feedID)
		if err != nil {
			return time.Time{}, err
		}

		tmp, err := json.Marshal(lastNewsTime)
		if err != nil {
			log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
			return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
		}

		if err := RedisClient.Set(ctx, key, tmp, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
			return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
		}

		return GetLastNewsTime(ctx, feedID)
	} else if err != nil {
		log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
		return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
	}

	if err := json.Unmarshal(lastNewsTimeByte, &lastNewsTime); err != nil {
		log.Sugar.Errorw("cache.GetLastNewsTime", "error", err)
		return time.Time{}, fmt.Errorf("cache.GetLastNewsTime: %w", err)
	}

	return lastNewsTime, nil
}

func SetLastNewsTime(ctx context.Context, feedID int) error {
	tmp, err := json.Marshal(time.Now())
	if err != nil {
		log.Sugar.Errorw("cache.SetLastNewsTime", "error", err)
		return fmt.Errorf("cache.SetLastNewsTime: %w", err)
	}

	key := fmt.Sprintf("lastNewsTime:%v", feedID)

	if err := RedisClient.Set(ctx, key, tmp, time.Hour).Err(); err != nil {
		log.Sugar.Errorw("cache.SetLastNewsTime", "error", err)
		return fmt.Errorf("cache.SetLastNewsTime: %w", err)
	}

	return nil
}
