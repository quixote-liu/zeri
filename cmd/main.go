package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"zeri/internal/config"
)

func main() {
	if err := config.CONF.LoadConfiguration("config.conf"); err != nil {
		log.Printf("load configuration config.conf failed: %v", err)
		return
	}

	r := gin.Default()

	r.Run(":8080")
}
