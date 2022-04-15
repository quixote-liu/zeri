package service

import (
	"errors"

	"zeri/internal/model/system"
	"zeri/internal/model/system/request"
	"zeri/internal/respository"
	"zeri/pkg/utils"

	"gorm.io/gorm"
)

type UserService interface {
	Login(u *request.Login) (userInter *system.SysUser, err error)
}

type userService struct {
	userRepo respository.UserRepository
}

func NewUserService(userRepo respository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (svc *userService) Login(u *request.Login) (userInter *system.SysUser, err error) {
	u.Password = utils.MD5V([]byte(u.Password))

	user, err := svc.userRepo.FindUserByNamePass(u.Username, u.Password)
	if err == nil {
		var am system.SysMenu
		ferr := global.GVA_DB.First(&am, "name = ? AND authority_id = ?", user.Authority.DefaultRouter, user.AuthorityId).Error
		if errors.Is(ferr, gorm.ErrRecordNotFound) {
			user.Authority.DefaultRouter = "404"
		}
	}
	return err, &user
}
