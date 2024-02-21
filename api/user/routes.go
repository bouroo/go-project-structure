package user

import (
	"net/http"

	"github.com/bouroo/go-project-structure/api/middleware"
	"github.com/bouroo/go-project-structure/api/user/handler"
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) *echo.Echo {
	router := e.Group("/api/v1/users")
	router.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, model.GeneralResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "Welcome to user API",
		})
	})

	router.GET("/me", handler.ReadUserDetails, middleware.CustomJWTMiddleware(datasources.AppConfig.GetString("jwt.key")))

	return e
}
