package handler

import (
	"log/slog"
	"net/http"

	"github.com/bouroo/go-clean-arch/domain"
	"github.com/bouroo/go-clean-arch/model"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type authHandler struct {
	config      *viper.Viper
	logger      *slog.Logger
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(config *viper.Viper, logger *slog.Logger, authUsecase domain.AuthUsecase) domain.AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
		config:      config,
		logger:      logger,
	}
}

func (h *authHandler) RegisterRoute(e *echo.Echo) *echo.Echo {
	router := e.Group("/api/v1/auth")

	router.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, model.GeneralResponse{
			Code:   http.StatusOK,
			Status: "success",
		})
	})

	router.POST("/signin", h.SignIn)
	return e
}
