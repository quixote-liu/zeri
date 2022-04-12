package cache

import (
	"context"
	"zeri/internal/config"

	"github.com/go-redis/redis/v8"
)

var conf = config.CONF

var rdb *redis.Client

func InitRedis() error {
	addr := conf.GetString("redis", "addr")
	pd := conf.GetString("redis", "password")
	db := conf.GetInt("redis", "db")

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pd,
		DB:       db,
	})

	if err := c.Ping(context.TODO()).Err(); err != nil {
		return err
	}

	rdb = c
	return nil
}
