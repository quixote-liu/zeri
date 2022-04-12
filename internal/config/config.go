package config

import "github.com/quixote-liu/config"

func init() {
	serverConfig()
}

var CONF = config.CONF

func serverConfig() {
	group := config.NewGroup("server")

	group.SetString("host", "127.0.0.1")
	group.SetString("port", "8080")
	CONF.RegisterGroup(group)
}
