package handler

import (
	"log/slog"
	"net/http"

	"github.com/bouroo/go-clean-arch/app/user"
	"github.com/bouroo/go-clean-arch/entity"
	"github.com/bouroo/go-clean-arch/model"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase user.Usecase
	Logger      *slog.Logger
}

func NewUserHandler(userUsecase user.Usecase, logger *slog.Logger) user.Handler {
	return &userHandler{
		userUsecase: userUsecase,
		Logger:      logger,
	}
}

func (u *userHandler) CreateUser(c echo.Context) (err error) {
	return
}

func (u *userHandler) GetUserDetails(c echo.Context) (err error) {
	response := model.GeneralResponse{
		Code:   http.StatusBadRequest,
		Status: "fail",
	}

	userID := c.Get("user_id").(string)

	userDetails, err := u.userUsecase.ReadUserDetails(userID, "", "")
	if err != nil {
		response.Message = err.Error()
		return c.JSON(response.Code, response)
	}

	response.Data = map[string]interface{}{
		"user_details": userDetails,
	}

	if err == nil {
		response.Code = http.StatusOK
		response.Status = "success"
	}
	return c.JSON(response.Code, response)
}

func (u *userHandler) UpdateProfile(c echo.Context) (err error) {
	response := model.GeneralResponse{
		Code:   http.StatusBadRequest,
		Status: "fail",
	}

	userID := c.Get("user_id").(string)
	var userProfile model.UserProfile
	err = c.Bind(&userProfile)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(response.Code, response)
	}

	err = u.userUsecase.UpdateUserAccount(userID, entity.UserAccount{
		UserProfile: entity.UserProfile{
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
			Email:     userProfile.Email,
			Phone:     userProfile.Phone,
			Avatar:    userProfile.Avatar,
		},
	})

	if err == nil {
		response.Code = http.StatusOK
		response.Status = "success"
	}
	return c.JSON(response.Code, response)
}
