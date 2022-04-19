package jwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	jwt.RegisteredClaims
	BufferTime int
	BaseClaims
}

type BaseClaims struct {
	ID          uint
	UUID        string
	UserName    string
	NickName    string
	AuthorityID string
}
