package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CustomRequestIDGenerator() string {
	return uuid.NewString()
}

func CustomRequestIDHandler(c echo.Context, rid string) {
	c.Set("request_id", rid)
}
