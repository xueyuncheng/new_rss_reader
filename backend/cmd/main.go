package main

import (
	"flag"
	"myNewFeed/cache"
	"myNewFeed/crontab"
	"myNewFeed/database"
	"myNewFeed/http"
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"myNewFeed/service"

	"github.com/BurntSushi/toml"
)

func main() {
	initPprof()

	file := flag.String("f", "./config.toml", "config file path")

	flag.Parse()

	config := &model.Config{}
	if _, err := toml.DecodeFile(*file, &config); err != nil {
		log.Sugar.Fatalw("配置文件读取失败", "err", err)
	}

	log.InitLog()
	database.InitDB(&config.Mysql)
	cache.InitRedis(&config.Redis)
	service.InitService(&config.Proxy)
	crontab.InitCrontab()
	http.InitHttp(&config.Http)
}
