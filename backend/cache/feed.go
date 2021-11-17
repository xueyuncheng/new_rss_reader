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

// 获取所有的rss源名称
func GetFeed(ctx context.Context) ([]*model.Feed, error) {
	var err error
	feeds := make([]*model.Feed, 0, 16)

	feedsString, err := RedisClient.Get(ctx, "feed").Result()
	if err == redis.Nil {
		feeds, err = database.GetFeed(ctx)
		if err != nil {
			return nil, err
		}

		feedsByte, err := json.Marshal(feeds)
		if err != nil {
			return nil, fmt.Errorf("json 序列化错误: %w", err)
		}

		if err := RedisClient.Set(ctx, "feed", feedsByte, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("redis set error", "err", err)
			return nil, fmt.Errorf("redis set error: %w", err)
		}

		return GetFeed(ctx)
	} else if err != nil {
		return nil, fmt.Errorf("redis get error: %v", err)
	}

	if err := json.Unmarshal([]byte(feedsString), &feeds); err != nil {
		log.Sugar.Errorw("json 反序列化错误", "err", err)
		return nil, fmt.Errorf("json 反序列化错误: %w", err)
	}

	return feeds, nil
}

func DeleteFeed(ctx context.Context) error {
	if err := RedisClient.Del(ctx, "feed").Err(); err != nil {
		return fmt.Errorf("redis del error: %w", err)
	}

	return nil
}
