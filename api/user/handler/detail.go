package handler

import (
	"net/http"

	"github.com/bouroo/go-project-structure/api/user/usecase"
	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

func ReadUserDetails(c echo.Context) (err error) {
	reqpPayload := model.GeneralResponse{
		Code:   http.StatusInternalServerError,
		Status: "error",
	}

	userID := c.Get("userID").(string)

	resp, err := usecase.ReadUserDetails(model.UserAccount{ID: userID})
	if err != nil {
		reqpPayload.Code = http.StatusInternalServerError
		reqpPayload.Status = "error"
		reqpPayload.Message = err.Error()
		return c.JSON(reqpPayload.Code, reqpPayload)
	}

	reqpPayload.Data = resp

	if err == nil {
		reqpPayload.Code = http.StatusOK
		reqpPayload.Status = "success"
	}
	return c.JSON(reqpPayload.Code, reqpPayload)
}
