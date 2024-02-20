package handler

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authUsecase domain.AuthUsecase
	Logger      *slog.Logger
}

func NewAuthHandler(authUsecase domain.AuthUsecase, logger *slog.Logger) domain.AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
		Logger:      logger,
	}
}

func (h *authHandler) RegisterRoute(e *echo.Echo) *echo.Echo {
	router := e.Group("/api/v1/auth")
	router.POST("/signin", h.SignIn)
	return e
}
