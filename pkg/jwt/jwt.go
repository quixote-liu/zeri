package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that is not even a token")
	TokenInvalid     = errors.New("cloud not handle this token")
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

func (c *client) CreateToken(claim BaseClaims) (string, error) {
	clm := CustomClaims{
		BaseClaims:       claim,
		RegisteredClaims: c.registeredClaims(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)
	return token.SignedString(c.signingKey)
}

func (c *client) registeredClaims() jwt.RegisteredClaims {
	expires := config.CONF.GetString("jwt", "expires_time")
	e, err := time.ParseDuration(expires)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("parse jwt expires time failed, will select default expires time(12 hours)")
		e = 12 * time.Hour
	}

	issuer := config.CONF.GetString("jwt", "issuer")

	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(e)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    issuer,
	}
}

func (c *client) ParseToken(tokenString string) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return c.signingKey, nil
	})
	if token.Valid {
		cc := CustomClaims{}
	}
	if err != nil {
		switch err.(type) {

		}
	}
}
