package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type standardClaims struct {
	Id   string `json:"id"`
	Name string `json:"username"`
	jwt.StandardClaims
}

func NewStandardClaims(id, name string, expiry time.Duration) jwt.Claims {
	return standardClaims{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Unix(),
		},
	}
}
