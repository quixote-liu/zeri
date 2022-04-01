package database

import "fmt"

type mysqlOptions struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}

var mysqlOptsInstance = func() mysqlOptions {
	group := "database"

	return mysqlOptions{
		username: CONF.GetString(group, "user_name"),
		password: CONF.GetString(group, "password"),
		host:     CONF.GetString(group, "host"),
		port:     CONF.GetString(group, "port"),
		dbName:   CONF.GetString(group, "db_name"),
	}
}()

func (opts mysqlOptions) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		opts.username, opts.password, opts.host, opts.port, opts.dbName)
}
