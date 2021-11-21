package cache

import (
	"fmt"
	"myNewFeed/model"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(config *model.Redis) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
		DB:       0, // use default DB
	})

	RedisClient = rdb
}
