package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
)

type Client interface {
}

type client struct {
	signingKey []byte
}

func New() Client {
	signingKey := config.CONF.GetString("jwt", "signing_key")

	return &client{
		signingKey: []byte(signingKey),
	}
}

func (c *client) standardClaims() jwt.StandardClaims {
	expires := config.CONF.GetString("jwt", "expires_time")
	e, err := time.ParseDuration(expires)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("parse jwt expires time failed, will select default expires time(12 hours)")
		e = 12 * time.Hour
	}

	issuer := config.CONF.GetString("jwt", "issuer")

	return jwt.StandardClaims{
		ExpiresAt: time.Now().Add(e).Unix(),
		Issuer:    issuer,
	}
}
