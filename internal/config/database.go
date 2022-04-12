package config

import "github.com/quixote-liu/config"

func init() {
	dataBaseConfig()
}

func dataBaseConfig() {
	group := config.NewGroup("batabase")

	group.SetString("user_name", "testing")
	group.SetString("password", "testing")
	group.SetString("host", "127.0.0.1")
	group.SetString("port", "8080")
	group.SetString("db_name", "zeri")
	group.SetString("db_type", "mysql")
	CONF.RegisterGroup(group)
}
