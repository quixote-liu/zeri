package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrfooRedisNil = redis.Nil

type Client interface {
	Set(key string, value interface{}) error
	SetWithTime(key string, value interface{}, expiration time.Duration) error
	GetString(key string) string
	GetInt(key string) (int, error)
}

type client struct {
	rdb *redis.Client
}

func New() Client {
	return &client{rdb: rdb}
}

func (c *client) Set(key string, value interface{}) error {
	return c.rdb.Set(context.TODO(), key, value, 0).Err()
}

func (c *client) SetWithTime(key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(context.TODO(), key, value, expiration).Err()
}

func (c *client) GetString(key string) string {
	return c.rdb.Get(context.TODO(), key).String()
}

func (c *client) GetInt(key string) (int, error) {
	return c.rdb.Get(context.TODO(), key).Int()
}
