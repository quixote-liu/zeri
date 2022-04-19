package repository

import (
	"zeri/internal/model/system"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByNamePass(userName, password string) (system.SysUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (rep *userRepository) FindUserByNamePass(userName, password string) (system.SysUser, error) {
	u := system.SysUser{}
	err := rep.db.Where("username = ? AND password = ?", u.Username, u.Password).
		Preload("Authorities").Preload("Authority").First(&u).Error
	return u, err
}
