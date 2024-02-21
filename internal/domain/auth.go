package domain

import "github.com/labstack/echo/v4"

type AuthRepository interface {
}

type AuthUsecase interface {
}

type AuthHandler interface {
	RegisterRoute(e *echo.Echo) *echo.Echo
}
