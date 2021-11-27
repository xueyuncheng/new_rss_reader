package crontab

import (
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"myNewFeed/service"

	"github.com/robfig/cron/v3"
)

func InitCrontab(config *model.CronTab) {
	c := cron.New()
	if _, err := c.AddFunc(config.Schedule, func() {
		service.RefreshNews()
	}); err != nil {
		log.Sugar.Fatalw("添加定时任务错误", "err", err)
	}

	c.Start()
}
