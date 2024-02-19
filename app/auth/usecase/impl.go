package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/auth"
)

type authUsecase struct {
	authRepo auth.Repository
	Logger   *slog.Logger
}

func NewAuthUsecase(authRepo auth.Repository, logger *slog.Logger) auth.Usecase {
	return &authUsecase{
		authRepo: authRepo,
		Logger:   logger,
	}
}
