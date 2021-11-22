package crontab

import (
	"myNewFeed/internal/log"
	"myNewFeed/service"

	"github.com/robfig/cron/v3"
)

func InitCrontab() {
	c := cron.New()
	if _, err := c.AddFunc("*/5 * * * *", func() {
		service.RefreshNews()
	}); err != nil {
		log.Sugar.Fatalw("添加定时任务错误", "err", err)
	}

	c.Start()
}
