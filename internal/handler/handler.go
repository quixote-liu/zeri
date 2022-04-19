package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
	"zeri/internal/model/common/response"
	"zeri/internal/model/system/request"
	sysResp "zeri/internal/model/system/response"
	"zeri/internal/service"
	"zeri/pkg/utils"
)

type BaseHandler interface {
	Login(c *gin.Context)
	Captcha(c *gin.Context)
}

type baseHandler struct {
	store   base64Captcha.Store
	userSvc service.UserService
	jwtSvc  service.JWTService
}

func NewBaseHandler(userSvc service.UserService, jwtSvc service.JWTService) BaseHandler {
	return &baseHandler{
		store:   base64Captcha.DefaultMemStore,
		userSvc: userSvc,
		jwtSvc:  jwtSvc,
	}
}

func (h *baseHandler) Login(c *gin.Context) {
	var l request.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证验证码
	if !h.store.Verify(l.CaptchaId, l.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}

	u, err := h.userSvc.Login(l)
	if err != nil {
		response.FailWithMessage("用户名或者密码错误", c)
		return
	}

	// 对用户发布token
	token, err := h.jwtSvc.IssueToken(u)
	if err != nil {
		response.FailWithMessage("设置登录状态失败", c)
		return
	}
	response.OkWithDetailed(sysResp.LoginResponse{
		User:      u,
		Token:     token,
		ExpiresAt: h.getExpiresAt(),
	}, "登录成功", c)
}

func (h *baseHandler) getExpiresAt() float64 {
	expires := config.CONF.GetString("jwt", "expires_time")
	e, err := time.ParseDuration(expires)
	if err != nil {
		log.Errorf("parse jwt expires time failed: %v", err)
		return 0
	}
	return e.Seconds()
}

func (h *baseHandler) Captcha(c *gin.Context) {
	imageHeight := config.CONF.GetInt("captcha", "img-height")
	imageWidth := config.CONF.GetInt("captcha", "img-width")
	keyLong := config.CONF.GetInt("captcha", "key-long")

	driver := base64Captcha.NewDriverDigit(imageHeight, imageWidth, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, h.store)
	if id, b64s, err := cp.Generate(); err != nil {
		log.Errorf("generate captcha failed: %v", err)
		utils.FailWithMessage("验证码获取失败", c)
	} else {
		utils.OkWithDetailed(sysResp.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: keyLong,
		}, "验证码获取成功", c)
	}
}
