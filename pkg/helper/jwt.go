package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(jwtKey string, claims *jwt.RegisteredClaims) (tokenStr string, err error) {
	if len(jwtKey) == 0 {
		return "", errors.New("empty jwtKey")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token == nil {
		return "", errors.New("failed to generate token")
	}
	return token.SignedString([]byte(jwtKey))
}

func ParseJWTToken(tokenStr string, jwtKey string) (claims *jwt.RegisteredClaims, err error) {
	if len(tokenStr) == 0 {
		return nil, errors.New("empty tokenStr")
	}
	if len(jwtKey) == 0 {
		return nil, errors.New("empty jwtKey")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims = token.Claims.(*jwt.RegisteredClaims)
	return
}
