package repository

import (
	"zeri/internal/model/system"

	"gorm.io/gorm"
)

type MenuRepository interface {
	FindMenuByNameAndAuthID(name, authID string) (system.SysMenu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMuneRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (repo *menuRepository) FindMenuByNameAndAuthID(name, authID string) (system.SysMenu, error) {
	menu := system.SysMenu{}
	err := repo.db.First(&menu, "name = ? AND authority_id = ?", name, authID).Error
	return menu, err
}
