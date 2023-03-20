package redis

import (
	"fmt"
	"gin-cli/src/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port,
		),
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})
	return nil
}

func Close() {
	rdb.Close()
}
