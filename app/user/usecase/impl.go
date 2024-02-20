package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
	Logger   *slog.Logger
}

func NewUserUsecase(userRepo domain.UserRepository, logger *slog.Logger) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		Logger:   logger,
	}
}
