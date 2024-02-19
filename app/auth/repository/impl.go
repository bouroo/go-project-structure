package repository

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/auth"
	"gorm.io/gorm"
)

type authRepository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewAuthRepository(db *gorm.DB, logger *slog.Logger) auth.Repository {
	return &authRepository{
		DB:     db,
		Logger: logger,
	}
}
