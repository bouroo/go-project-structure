package helper

import (
	"errors"
	"net/http"

	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

// CustomHTTPErrorHandler is a Go function that handles HTTP errors.
//
// It takes err error and ctx echo.Context as parameters.
func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	var httpStatus = http.StatusInternalServerError
	var resp model.Response
	var echoError *echo.HTTPError
	var customError *model.Error

	if errors.As(err, &echoError) {
		httpStatus = echoError.Code
		if httpStatus >= 400 && httpStatus < 500 {
			resp.Status = model.RespFail
		} else {
			resp.Status = model.RespError
		}
		resp.Code = echoError.Code
	} else if errors.As(err, &customError) {
		if customError.HTTPStatus != 0 {
			httpStatus = customError.HTTPStatus
		}
		if customError.Data != nil {
			resp.Data = customError.Data
		}
		resp.Code = customError.Code
		resp.Status = customError.Status
	}
	resp.Message = err.Error()
	ctx.Logger().Error(err)
	if err = ctx.JSON(httpStatus, resp); err != nil {
		ctx.Logger().Error(err)
	}
}
