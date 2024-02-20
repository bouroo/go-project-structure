package handler

import (
	"net/http"

	"github.com/bouroo/go-clean-arch/model"
	"github.com/labstack/echo/v4"
)

func (h *authHandler) SignIn(c echo.Context) (err error) {
	var respPayload model.TokenResponse
	var reqPayload model.Signin

	err = c.Bind(&reqPayload)
	if err != nil {
		respPayload = model.TokenResponse{
			Error: model.OauthErrInvalidRequest,
		}
		return c.JSON(http.StatusBadRequest, respPayload)
	}

	// TODO: Validate request

	return c.JSON(http.StatusOK, respPayload)
}
