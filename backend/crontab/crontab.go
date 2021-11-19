package crontab

import (
	"myNewFeed/internal/log"
	"myNewFeed/service"

	"github.com/robfig/cron/v3"
)

func InitCrontab() {
	c := cron.New()
	if _, err := c.AddFunc("* * * * *", func() {
		service.RefreshNews()
	}); err != nil {
		log.Sugar.Fatalw("init crontab failed", "err", err)
	}

	c.Start()
}
