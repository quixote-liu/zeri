package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"zeri/internal/config"
)

var CONF = config.CONF

func main() {
	if err := CONF.LoadConfiguration("config.conf"); err != nil {
		log.Printf("load configuration config.conf failed: %v", err)
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
