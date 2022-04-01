package database

import (
	"fmt"
	"zeri/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var CONF = config.CONF

var DB *gorm.DB

func InitDataBase() error {
	var dsn string

	dbtype := CONF.GetString("database", "db_type")
	switch dbtype {
	case "mysql":
		dsn = mysqlOptsInstance.dsn()
	default:
		return fmt.Errorf("the does not support database type %s, only support `mysql`", dbtype)
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(); err != nil {
		return err
	}

	return nil
}
