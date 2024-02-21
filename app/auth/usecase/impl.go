package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/domain"
	"github.com/spf13/viper"
)

type authUsecase struct {
	config   *viper.Viper
	logger   *slog.Logger
	authRepo domain.AuthRepository
	userRepo domain.UserRepository
}

func NewAuthUsecase(config *viper.Viper, logger *slog.Logger, authRepo domain.AuthRepository, userRepo domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		config:   config,
		logger:   logger,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
