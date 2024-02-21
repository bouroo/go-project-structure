package auth

import (
	"net/http"

	"github.com/bouroo/go-project-structure/api/auth/handler"
	"github.com/bouroo/go-project-structure/api/middleware"
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) *echo.Echo {
	router := e.Group("/api/v1/auth")
	router.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, model.GeneralResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "Welcome to Auth API",
		})
	})

	router.POST("/signup", handler.Signup)
	router.POST("/signin", handler.Signin)
	router.POST("/change-password", handler.ChangePassword, middleware.CustomJWTMiddleware(datasources.AppConfig.GetString("jwt.key")))
	return e
}
