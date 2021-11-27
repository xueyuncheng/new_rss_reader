package cache

import (
	"myNewFeed/model"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis(config *model.Redis) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
	})
}
