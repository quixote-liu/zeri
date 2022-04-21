package database

import (
	"fmt"
	"zeri/internal/config"
	"zeri/internal/model/example"
	"zeri/internal/model/system"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() error {
	var err error

	dbtype := config.CONF.GetString("database", "db_type")

	switch dbtype {
	case "mysql":
		DB, err = gormMysql()
	default:
		return fmt.Errorf("the does not support database type %s, only support `mysql`", dbtype)
	}
	if err != nil {
		return fmt.Errorf("init database failed: %v", err)
	}

	return registerTables(DB)
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
