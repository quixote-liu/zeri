package config

import "github.com/quixote-liu/config"

func init() {
	serverConfig()
}

var CONF = config.CONF

func serverConfig() {
	group := config.NewGroup("system")

	group.SetString("host", "127.0.0.1")
	group.SetString("port", "8888")
	group.SetString("db_type", "mysql")
	group.SetBool("multipoint", false)
	CONF.RegisterGroup(group)
}
