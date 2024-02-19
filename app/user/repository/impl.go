package repository

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/user"
	"gorm.io/gorm"
)

type userRepository struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func NewUserRepository(db *gorm.DB, logger *slog.Logger) user.Repository {
	return &userRepository{
		DB:     db,
		Logger: logger,
	}
}
