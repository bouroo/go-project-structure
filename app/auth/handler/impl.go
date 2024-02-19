package handler

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/auth"
)

type authHandler struct {
	authUsecase auth.Usecase
	Logger      *slog.Logger
}

func NewAuthHandler(authUsecase auth.Usecase, logger *slog.Logger) auth.Handler {
	return &authHandler{
		authUsecase: authUsecase,
		Logger:      logger,
	}
}
