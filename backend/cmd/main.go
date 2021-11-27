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
	"os"

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

	os.Setenv("HTTP_PROXY", config.Proxy.Address)

	log.InitLog()
	database.InitDB(&config.Mysql)
	cache.InitRedis(&config.Redis)
	service.InitService()
	crontab.InitCrontab(&config.CronTab)
	http.InitHttp(&config.Http)
}
