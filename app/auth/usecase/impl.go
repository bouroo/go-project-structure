package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/domain"
)

type authUsecase struct {
	authRepo domain.AuthRepository
	userRepo domain.UserRepository
	Logger   *slog.Logger
}

func NewAuthUsecase(authRepo domain.AuthRepository, userRepo domain.UserRepository, logger *slog.Logger) domain.AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
		userRepo: userRepo,
		Logger:   logger,
	}
}
