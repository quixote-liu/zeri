package service

import (
	"zeri/internal/model/system"
	"zeri/internal/model/system/request"
	"zeri/internal/repository"
	"zeri/pkg/utils"
)

type UserService interface {
	Login(u request.Login) (userInter system.SysUser, err error)
}

type userService struct {
	userRepo repository.UserRepository
	menuRepo repository.MenuRepository
}

func NewUserService(userRepo repository.UserRepository,
	menuRepo repository.MenuRepository) UserService {
	return &userService{
		userRepo: userRepo,
		menuRepo: menuRepo,
	}
}

func (svc *userService) Login(u request.Login) (userInter system.SysUser, err error) {
	u.Password = utils.MD5V([]byte(u.Password))

	user, err := svc.userRepo.FindUserByNamePass(u.Username, u.Password)
	if err != nil {
		return
	}

	_, err = svc.menuRepo.FindMenuByNameAndAuthID(user.Authority.DefaultRouter, user.AuthorityId)
	if err != nil {
		user.Authority.DefaultRouter = "404"
	}
	return user, nil
}
