package handler

import (
	"net/http"

	"github.com/bouroo/go-project-structure/api/auth/usecase"
	"github.com/bouroo/go-project-structure/pkg/model"
	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) (err error) {
	respPayload := model.GeneralResponse{
		Code:   http.StatusInternalServerError,
		Status: "error",
	}
	var userAccount model.UserAccount
	err = c.Bind(&userAccount)
	if err != nil {
		respPayload.Code = http.StatusBadRequest
		respPayload.Status = "fail"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}
	resp, err := usecase.Singup(userAccount.Email, userAccount.Password)
	if err != nil {
		respPayload.Code = http.StatusInternalServerError
		respPayload.Status = "error"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}

	respPayload.Data = resp

	if err == nil {
		respPayload.Code = http.StatusOK
		respPayload.Status = "success"
		respPayload.Message = "User created successfully"
	}
	return c.JSON(respPayload.Code, respPayload)
}

func ChangePassword(c echo.Context) (err error) {
	respPayload := model.GeneralResponse{
		Code:   http.StatusInternalServerError,
		Status: "error",
	}

	var userAccount model.UserAccount
	err = c.Bind(&userAccount)
	if err != nil {
		respPayload.Code = http.StatusBadRequest
		respPayload.Status = "fail"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}

	userAccount.ID = c.Get("userID").(string)

	err = usecase.ChangePassword(userAccount.ID, userAccount.Password)
	if err != nil {
		respPayload.Code = http.StatusInternalServerError
		respPayload.Status = "error"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}

	if err == nil {
		respPayload.Code = http.StatusOK
		respPayload.Status = "success"
		respPayload.Message = "Password changed successfully"
	}
	return c.JSON(respPayload.Code, respPayload)
}

func Signin(c echo.Context) (err error) {
	respPayload := model.GeneralResponse{
		Code:   http.StatusInternalServerError,
		Status: "error",
	}

	var userAccount model.UserAccount
	err = c.Bind(&userAccount)
	if err != nil {
		respPayload.Code = http.StatusBadRequest
		respPayload.Status = "fail"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}

	token, err := usecase.Signin(userAccount.Email, userAccount.Password)
	if err != nil {
		respPayload.Code = http.StatusInternalServerError
		respPayload.Status = "error"
		respPayload.Message = err.Error()
		return c.JSON(respPayload.Code, respPayload)
	}

	respPayload.Data = map[string]string{
		"token": token,
	}
	if err == nil {
		respPayload.Code = http.StatusOK
		respPayload.Status = "success"
		respPayload.Message = "User logged in successfully"
	}

	return c.JSON(respPayload.Code, respPayload)
}
