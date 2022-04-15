package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
	"zeri/pkg/utils"
)

func Login(c *gin.Context) {

}

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	imageHeight := config.CONF.GetInt("captcha", "img-height")
	imageWidth := config.CONF.GetInt("captcha", "img-width")
	imageLong := config.CONF.GetInt("captcha", "key-long")

	driver := base64Captcha.NewDriverDigit(imageHeight, imageWidth, imageLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		log.Errorf("generate captcha failed: %v", err)
		utils.FailWithMessage("验证码获取失败", c)
	} else {
		utils.OkWithDetailed()
	}
}
