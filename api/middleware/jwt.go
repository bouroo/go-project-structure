package middleware

import (
	"net/http"
	"strings"

	"github.com/bouroo/go-project-structure/pkg/helper"
	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

func CustomJWTMiddleware(jwtKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			authPath := strings.Fields(authHeader)
			if len(authPath) != 2 || authPath[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, model.GeneralResponse{
					Code:    http.StatusUnauthorized,
					Status:  "fail",
					Message: "unauthorized access, please provide a valid token",
				})
			}
			claims, err := helper.ParseJWTToken(authPath[1], jwtKey)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, model.GeneralResponse{
					Code:    http.StatusUnauthorized,
					Status:  "fail",
					Message: err.Error(),
				})
			}
			c.Set("user_id", claims.Subject)
			return next(c)
		}
	}
}
