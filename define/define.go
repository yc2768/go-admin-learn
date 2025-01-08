package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	JwtKey             = []byte("asd123456789")
	TokenExpire        = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
)

type UserClaims struct {
	Id      uint
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}
