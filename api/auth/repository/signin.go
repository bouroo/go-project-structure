package repository

import (
	"time"

	"github.com/bouroo/go-project-structure/datasources"
	"github.com/golang-jwt/jwt/v5"
)

func GenAccessToken(userID, email string) (tokenString string, err error) {

	exp := jwt.NewNumericDate(time.Now().Add(datasources.AppConfig.GetDuration("jwt.ttl")))

	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    email,
		ExpiresAt: exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(datasources.AppConfig.GetString("jwt.key")))

	return
}