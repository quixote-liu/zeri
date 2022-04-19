package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zeri/pkg/cache"
	"zeri/pkg/database"
	"zeri/pkg/jwt"
)

func Router() *gin.Engine {
	r := gin.Default()
	cache := cache.New()
	db := database.DB
	jwt := jwt.New()

	// 健康检测
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 静态文件
	r.Static("/form-generator", "./resource/page")

	// 基础路由
	{
		h := InitializeBaseHandler(db, cache, jwt)
		baseAPI := r.Group("/base")
		baseAPI.POST("/login", h.Login)
		baseAPI.POST("/captcha", h.Captcha)
	}

	return r
}
