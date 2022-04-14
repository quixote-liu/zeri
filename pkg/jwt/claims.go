package jwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	BufferTime  string
	UUID        string
	UserName    string
	NickName    string
	AuthorityID string
}
