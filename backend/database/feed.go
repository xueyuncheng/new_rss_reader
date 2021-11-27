package database

import (
	"context"
	"fmt"
	"myNewFeed/internal/log"
	"myNewFeed/model"
)

func GetFeed(ctx context.Context) ([]*model.Feed, error) {
	feeds := make([]*model.Feed, 0, 16)
	if err := db.Find(&feeds).Error; err != nil {
		log.Sugar.Errorw("获取rss源错误", "err", err)
		return nil, fmt.Errorf("获取rss源错误: %v", err)
	}

	return feeds, nil
}

func AddFeed(ctx context.Context, feed *model.Feed) error {
	if err := db.Create(feed).Error; err != nil {
		log.Sugar.Errorw("添加rss源错误", "err", err)
		return fmt.Errorf("添加rss源错误: %v", err)
	}

	return nil
}

func DeleteFeed(ctx context.Context, id int) error {
	if err := db.Where("id = ?", id).Delete(&model.Feed{}).Error; err != nil {
		log.Sugar.Errorw("删除rss源错误", "err", err)
		return fmt.Errorf("删除rss源错误: %v", err)
	}

	return nil
}

func ListFeed(ctx context.Context) ([]*model.Feed, error) {
	feeds := make([]*model.Feed, 0, 16)
	if err := db.Find(&feeds).Error; err != nil {
		log.Sugar.Errorw("获取rss源错误", "err", err)
		return nil, fmt.Errorf("获取rss源错误: %v", err)
	}

	return feeds, nil
}
