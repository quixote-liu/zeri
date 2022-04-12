package main

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
	"zeri/pkg/cache"
	"zeri/pkg/database"
)

var CONF = config.CONF

func main() {
	// configuration
	if err := CONF.LoadConfiguration("config.conf"); err != nil {
		log.Errorf("load configuration config.conf failed: %v", err)
		return
	}

	// database
	if err := database.InitDataBase(); err != nil {
		log.Errorf("init database failed: %v", err)
		return
	}

	// redis
	if err := cache.InitRedis(); err != nil {
		log.Errorf("init redis failed: %v", err)
		return
	}

	r := gin.Default()

	host := CONF.GetString("server", "host")
	port := CONF.GetString("server", "port")
	addr := net.JoinHostPort(host, port)
	s := httpServer(r, addr)
	s.ListenAndServe()
}

func httpServer(h http.Handler, addr string) http.Server {
	return http.Server{
		Handler:           h,
		Addr:              addr,
		WriteTimeout:      20 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
	}
}
