package database

import (
	"fmt"
	"zeri/internal/config"
	"zeri/internal/model/example"
	"zeri/internal/model/system"

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
		dsn = mysqlOpts.dsn()
	default:
		return fmt.Errorf("the does not support database type %s, only support `mysql`", dbtype)
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := registerTables(DB); err != nil {
		return err
	}

	return nil
}

func registerTables(db *gorm.DB) error {
	tables := []interface{}{
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 示例模块表
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	}
	return db.AutoMigrate(tables...)
}
