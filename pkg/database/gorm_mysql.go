package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"zeri/internal/config"
)

func gormMysql() (*gorm.DB, error) {
	dsn := mysqlDSN()
	return gorm.Open(mysql.Open(dsn), gormConfig())
}

func mysqlDSN() string {
	group := "database"
	username := config.CONF.GetString(group, "user_name")
	password := config.CONF.GetString(group, "password")
	host := config.CONF.GetString(group, "host")
	port := config.CONF.GetString(group, "port")
	dbName := config.CONF.GetString(group, "db_name")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username, password, host, port, dbName)
}
