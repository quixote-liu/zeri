package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"

	"zeri/internal/config"
)

type Client interface {
	CreateToken(claim BaseClaims) (string, error)
	ParseToken(tokenString string) (*CustomClaims, error)
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
		BufferTime:       config.CONF.GetInt("jwt", "buffer_time"),
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

var (
	ErrfooTokenExpired     = errors.New("token is expired")
	ErrfooTokenNotValidYet = errors.New("token not active yet")
	ErrfooTokenMalformed   = errors.New("that is not even a token")
)

func (c *client) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return c.signingKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrfooTokenMalformed
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrfooTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrfooTokenNotValidYet
		}
		return nil, fmt.Errorf("clould not handle this token: %v", err)
	}

	cc, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("parse token string into custom claims failed")
	}
	return cc, nil
}
