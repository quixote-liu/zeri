package main

import (
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
	"zeri/internal/handler"
	"zeri/pkg/cache"
	"zeri/pkg/database"
)

var CONF = config.CONF

func main() {
	// parse configuration
	if err := CONF.LoadConfiguration("config.conf"); err != nil {
		log.Errorf("load configuration config.conf failed: %v", err)
		return
	}

	// set logrus.
	setupLogrus()

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

	r := handler.Router()

	host := CONF.GetString("system", "host")
	port := CONF.GetString("system", "port")
	addr := net.JoinHostPort(host, port)
	s := httpServer(r, addr)
	log.Infof("start zeri server on %s", addr)
	log.Error(s.ListenAndServe())
}

func httpServer(h http.Handler, addr string) http.Server {
	return http.Server{
		Handler:           h,
		Addr:              addr,
		WriteTimeout:      20 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
	}
}

func setupLogrus() {
	level := CONF.GetString("system", "log_level")
	l, err := log.ParseLevel(level)
	if err != nil {
		log.WithFields(log.Fields{
			"level": level,
			"error": err.Error(),
		}).Error("parse level of the configuration failed, will default set info level")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(l)
}
