package crontab

import (
	"myNewFeed/service"

	"github.com/robfig/cron/v3"
)

func InitCrontab() {
	c := cron.New()
	c.AddFunc("0 * * * * *", func() {
		service.RefreshNews()
	})

	c.Start()
}
