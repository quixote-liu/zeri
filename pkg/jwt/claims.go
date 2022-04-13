package jwt

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	jwt.StandardClaims
	BaseClaims
}

type BaseClaims struct {
	BufferTime  string
	UUID        string
	UserName    string
	NickName    string
	AuthorityID string
}
