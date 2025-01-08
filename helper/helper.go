package helper

import (
	"github.com/dgrijalva/jwt-go"
	"go-admin/define"
)

func GenerateToken(id uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaims{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
