package config

import "github.com/quixote-liu/config"

func init() {
	redisConfig()
}

func redisConfig() {
	g := config.NewGroup("redis")

	g.SetString("addr", "127.0.0.1:6379")
	g.SetString("password", "")
	g.SetInt("db", 0)
	CONF.RegisterGroup(g)
}
