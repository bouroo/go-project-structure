package repository

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/bouroo/go-clean-arch/infrastructure"
	"github.com/bouroo/go-clean-arch/internal/domain"
	"github.com/bouroo/go-clean-arch/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type authRepository struct {
	config    *viper.Viper
	logger    *slog.Logger
	db        *gorm.DB
	redisConn *infrastructure.RedisConn
}

func NewAuthRepository(config *viper.Viper, logger *slog.Logger, db *gorm.DB, redisConn *infrastructure.RedisConn) domain.AuthRepository {
	return &authRepository{
		config:    config,
		logger:    logger,
		db:        db,
		redisConn: redisConn,
	}
}

func (r *authRepository) GenerateToken(userID string, email string, jwtKey string) (token string, err error) {
	randBytes, _ := helper.GenerateRandomBytes(16)
	jti := helper.ComputeCRC64(randBytes)
	exp := time.Now().Add(r.config.GetDuration("jwt.ttl") * time.Second)

	claims := &jwt.RegisteredClaims{
		ID:        jti,
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token, err = helper.GenerateJWTToken(jwtKey, claims)
	if err != nil {
		return
	}

	redisKey := fmt.Sprintf("%s:%s", userID, jti)
	_, err = r.redisConn.Client.Set(r.redisConn.Context, redisKey, email, r.config.GetDuration("jwt.ttl")).Result()

	return
}

func (r *authRepository) ReadTokenRedis(userID string, jti string) (email string, err error) {
	redisKey := fmt.Sprintf("%s:%s", userID, jti)
	email, err = r.redisConn.Client.Get(r.redisConn.Context, redisKey).Result()
	return
}
