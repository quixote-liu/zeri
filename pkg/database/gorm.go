package database

import (
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func gormConfig() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   gormLogger(),
	}
}

func gormLogger() logger.Interface {
	c := logger.New(log.New(), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      false,
	})
	return c
}
