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
func ListFeed(ctx context.Context) ([]*model.Feed, error) {
	feedsBytes, err := rdb.Get(ctx, "feed").Bytes()
	if err == redis.Nil {
		feeds, err := database.GetFeed(ctx)
		if err != nil {
			return nil, err
		}

		feedsByte, err := json.Marshal(feeds)
		if err != nil {
			return nil, fmt.Errorf("json 序列化错误: %w", err)
		}

		if err := rdb.Set(ctx, "feed", feedsByte, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("redis set error", "err", err)
			return nil, fmt.Errorf("redis set error: %w", err)
		}

		return ListFeed(ctx)
	} else if err != nil {
		return nil, fmt.Errorf("redis get error: %v", err)
	}

	feeds := make([]*model.Feed, 0, 1024)
	if err := json.Unmarshal(feedsBytes, &feeds); err != nil {
		log.Sugar.Errorw("json 反序列化错误", "err", err)
		return nil, fmt.Errorf("json 反序列化错误: %w", err)
	}

	return feeds, nil
}

func DeleteFeed(ctx context.Context) error {
	if err := rdb.Del(ctx, "feed").Err(); err != nil {
		return fmt.Errorf("redis del error: %w", err)
	}

	return nil
}

func DeleteFeedByID(ctx context.Context, id int) error {
	key := fmt.Sprintf("feed:%v", id)
	if err := rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis del error: %w", err)
	}

	return nil
}
