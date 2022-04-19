//go:build wireinject

package handler

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"zeri/internal/repository"
	"zeri/internal/service"
	"zeri/pkg/cache"
	"zeri/pkg/jwt"
)

func InitializeBaseHandler(db *gorm.DB, cache cache.Client, jwt jwt.Client) BaseHandler {
	wire.Build(
		NewBaseHandler,
		service.NewUserService,
		service.NewJwtService,
		repository.NewUserRepository,
		repository.NewMuneRepository,
		repository.NewJWTRepository,
	)
	return nil
}
