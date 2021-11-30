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
			log.Sugar.Errorw("json marshal error", "err", err, "feeds", feeds)
			return nil, fmt.Errorf("json 序列化错误: %w", err)
		}

		if err := rdb.Set(ctx, "feed", feedsByte, time.Hour).Err(); err != nil {
			log.Sugar.Errorw("redis set error", "err", err, "key", "feed", "value", feedsByte)
			return nil, fmt.Errorf("redis set error: %w", err)
		}

		return ListFeed(ctx)
	} else if err != nil {
		log.Sugar.Errorw("redis get error", "err", err)
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
		log.Sugar.Errorw("redis del error", "err", err)
		return fmt.Errorf("redis del error: %w", err)
	}

	return nil
}

func DeleteFeedByID(ctx context.Context, id int) error {
	key := fmt.Sprintf("feed:%v", id)
	if err := rdb.Del(ctx, key).Err(); err != nil {
		log.Sugar.Errorw("redis del error", "err", err, "key", key)
		return fmt.Errorf("redis del error: %w", err)
	}

	return nil
}

func SetFeedStatus(ctx context.Context, id int, ok bool) error {
	key := fmt.Sprintf("status.feed:%v", id)
	if err := rdb.Set(ctx, key, ok, time.Hour).Err(); err != nil {
		log.Sugar.Errorw("redis set error", "err", err, "key", key, "value", ok)
		return fmt.Errorf("redis set error: %w", err)
	}

	return nil
}

func GetFeedStatus(ctx context.Context, id int) (bool, error) {
	key := fmt.Sprintf("status.feed:%v", id)
	ok, err := rdb.Get(ctx, key).Bool()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.Sugar.Errorw("redis get error", "err", err)
		return false, fmt.Errorf("redis get error: %w", err)
	}

	return ok, nil
}
