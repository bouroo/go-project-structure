package repository

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewAuthRepository(db *gorm.DB, logger *slog.Logger) domain.AuthRepository {
	return &authRepository{
		DB:     db,
		Logger: logger,
	}
}
