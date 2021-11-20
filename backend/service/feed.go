package service

import (
	"context"
	"myNewFeed/cache"
	"myNewFeed/database"
	"myNewFeed/model"
)

// 获取所有的 RSS 源
func ListFeed(ctx context.Context) ([]*model.Feed, error) {
	return cache.ListFeed(ctx)
}

// 添加一个 RSS 源
func AddFeed(ctx context.Context, req *model.AddFeedReq) error {
	feed := &model.Feed{
		Name: req.Name,
	}

	if err := database.AddFeed(ctx, feed); err != nil {
		return err
	}

	return cache.DeleteFeed(ctx)
}

func DeleteFeed(ctx context.Context, id int) error {
	if err := database.DeleteFeed(ctx, id); err != nil {
		return err
	}

	return cache.DeleteFeed(ctx)
}
