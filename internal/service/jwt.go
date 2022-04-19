package service

import (
	"fmt"
	"time"
	"zeri/internal/config"
	"zeri/internal/model/system"
	"zeri/internal/repository"
	"zeri/pkg/cache"
	"zeri/pkg/jwt"
)

type JWTService interface {
	IssueToken(user system.SysUser) (string, error)
}

type jwtService struct {
	cache   cache.Client
	jwt     jwt.Client
	jwtRepo repository.JWTRepository
}

func NewJwtService(
	cache cache.Client,
	jwt jwt.Client,
	jwtRepo repository.JWTRepository) JWTService {

	return &jwtService{
		cache:   cache,
		jwt:     jwt,
		jwtRepo: jwtRepo,
	}
}

func (svc *jwtService) IssueToken(user system.SysUser) (string, error) {
	token, err := svc.createToken(user)
	if err != nil {
		return "", err
	}

	if !svc.isMultipoint() {
		return token, nil
	}

	jwtVal := svc.getRedisJWT(user.Username)
	if jwtVal != "" {
		// 将旧的jwt值存入黑名单中既作废处理
		if err := svc.jwtRepo.CreateBlacklist(system.JwtBlacklist{Jwt: jwtVal}); err != nil {
			return "", fmt.Errorf("jwt作废失败: %v", err.Error())
		}
		if err := svc.setRedisBlackJWT(user.Username, jwtVal); err != nil {
			return "", fmt.Errorf("jwt作废失败: %v", err.Error())
		}
	}

	// 将新的jwt值存入缓存
	if err := svc.setRedisJWT(user.Username, token); err != nil {
		return "", err
	}

	return token, nil
}

func (svc *jwtService) createToken(user system.SysUser) (string, error) {
	claim := jwt.BaseClaims{
		ID:          user.ID,
		UUID:        user.UUID.String(),
		NickName:    user.NickName,
		UserName:    user.Username,
		AuthorityID: user.AuthorityId,
	}
	return svc.jwt.CreateToken(claim)
}

func (svc *jwtService) isMultipoint() bool {
	return config.CONF.GetBool("server", "multipoint")
}

func (svc *jwtService) getRedisJWT(key string) string {
	return svc.cache.GetString(key)
}

func (svc *jwtService) setRedisBlackJWT(key, jwtVal string) error {
	k := "black_list_" + key
	return svc.setRedisJWT(k, jwtVal)
}

func (svc *jwtService) setRedisJWT(key, jwtVal string) error {
	expire := config.CONF.GetString("jwt", "expires_time")
	t, err := time.ParseDuration(expire)
	if err != nil {
		return fmt.Errorf("parse jwt expires_time failed: %v", err)
	}
	return svc.cache.SetWithTime(key, jwtVal, t)
}
