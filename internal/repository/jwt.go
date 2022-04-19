package repository

import (
	"zeri/internal/model/system"

	"gorm.io/gorm"
)

type JWTRepository interface {
	CreateBlacklist(jwt system.JwtBlacklist) error
}

type jwtRepositroy struct {
	db *gorm.DB
}

func NewJWTRepository(db *gorm.DB) JWTRepository {
	return &jwtRepositroy{db: db}
}

func (repo *jwtRepositroy) CreateBlacklist(jwt system.JwtBlacklist) error {
	return repo.db.Create(&jwt).Error
}
