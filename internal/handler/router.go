package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 健康检测
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 静态文件
	r.Static("/form-generator", "./resource/page")

	return r
}
