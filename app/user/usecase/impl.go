package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/user"
)

type userUsecase struct {
	userRepo user.Repository
	Logger   *slog.Logger
}

func NewUserUsecase(userRepo user.Repository, logger *slog.Logger) user.Usecase {
	return &userUsecase{
		userRepo: userRepo,
		Logger:   logger,
	}
}
