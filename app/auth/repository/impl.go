package repository

import (
	"log/slog"
	"time"

	"github.com/bouroo/go-clean-arch/domain"
	"github.com/bouroo/go-clean-arch/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type authRepository struct {
	config *viper.Viper
	logger *slog.Logger
	db     *gorm.DB
}

func NewAuthRepository(config *viper.Viper, logger *slog.Logger, db *gorm.DB) domain.AuthRepository {
	return &authRepository{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (r *authRepository) GenerateToken(userID string, email string, jwtKey string) (token string, err error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(r.config.GetDuration("jwt.ttl") * time.Second)),
	}
	token, err = helper.GenerateJWTToken(jwtKey, claims)
	return
}
