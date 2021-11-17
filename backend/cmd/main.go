package main

import (
	"myNewFeed/cache"
	"myNewFeed/crontab"
	"myNewFeed/database"
	"myNewFeed/http"
	"myNewFeed/internal/log"
	"myNewFeed/service"
)

func main() {
	log.InitLog()
	database.InitDB()
	cache.InitRedis()
	service.InitService()
	crontab.InitCrontab()
	http.InitHttp()
}
