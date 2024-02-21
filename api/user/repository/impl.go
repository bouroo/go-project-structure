package repository

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/internal/domain"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type userRepository struct {
	config *viper.Viper
	logger *slog.Logger
	db     *gorm.DB
}

func NewUserRepository(config *viper.Viper, logger *slog.Logger, db *gorm.DB) domain.UserRepository {
	return &userRepository{
		config: config,
		logger: logger,
		db:     db,
	}
}
