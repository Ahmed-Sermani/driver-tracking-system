package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager struct {
	secretKey []byte
	expiry    time.Duration
}

func NewTokenManager(secretKey string, expiry time.Duration) *TokenManager {
	return &TokenManager{secretKey: []byte(secretKey), expiry: expiry}
}

func (tm *TokenManager) NewToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tm.secretKey)
}

func (tm *TokenManager) NewTokenWithStandardClaims(id, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewStandardClaims(id, name, tm.expiry))
	return token.SignedString(tm.secretKey)
}

func (tm *TokenManager) ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return tm.secretKey, nil
	})
	return err
}
