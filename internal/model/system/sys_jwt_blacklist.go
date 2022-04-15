package system

import "zeri/internal/model"

type JwtBlacklist struct {
	model.Base
	Jwt string `gorm:"type:text;comment:jwt"`
}
