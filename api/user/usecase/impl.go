package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/internal/domain"
	"github.com/spf13/viper"
)

type userUsecase struct {
	config   *viper.Viper
	logger   *slog.Logger
	userRepo domain.UserRepository
}

func NewUserUsecase(config *viper.Viper, logger *slog.Logger, userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		config:   config,
		logger:   logger,
		userRepo: userRepo,
	}
}
