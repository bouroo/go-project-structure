package handler

import (
	"log/slog"
	"net/http"

	"github.com/bouroo/go-clean-arch/internal/domain"
	"github.com/bouroo/go-clean-arch/internal/entity"
	"github.com/bouroo/go-clean-arch/internal/middleware"
	"github.com/bouroo/go-clean-arch/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type userHandler struct {
	config      *viper.Viper
	logger      *slog.Logger
	userUsecase domain.UserUsecase
}

func NewUserHandler(config *viper.Viper, logger *slog.Logger, userUsecase domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		config:      config,
		logger:      logger,
		userUsecase: userUsecase,
	}
}

func (h *userHandler) RegisterRoute(e *echo.Echo) *echo.Echo {
	router := e.Group("/api/v1/users")

	router.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, model.GeneralResponse{
			Code:   http.StatusOK,
			Status: "success",
		})
	})

	router.GET("/me", h.GetUserDetails, middleware.CustomJWTMiddleware(h.config.GetString("jwt.key")))

	router.POST("/register", h.CreateUser)

	return e
}

func (h *userHandler) CreateUser(c echo.Context) (err error) {
	response := model.GeneralResponse{
		Code:   http.StatusBadRequest,
		Status: "fail",
	}

	var user model.UserAccount
	err = c.Bind(&user)
	if err != nil {
		response.Message = err.Error()
		return c.JSON(response.Code, response)
	}

	err = h.userUsecase.CreateUserAccount(&entity.UserAccount{
		Email:    user.Email,
		Password: user.Password,
	})

	if err == nil {
		response.Code = http.StatusOK
		response.Status = "success"
	}
	return c.JSON(response.Code, response)
}

func (h *userHandler) GetUserDetails(c echo.Context) (err error) {
	response := model.GeneralResponse{
		Code:   http.StatusBadRequest,
		Status: "fail",
	}

	userID := c.Get("user_id").(string)

	userDetails, err := h.userUsecase.ReadUserDetails(userID, "")
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

func (h *userHandler) UpdateProfile(c echo.Context) (err error) {
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

	err = h.userUsecase.UpdateUserProfile(userID, entity.UserProfile{
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		Phone:     userProfile.Phone,
		Avatar:    userProfile.Avatar,
	})

	if err == nil {
		response.Code = http.StatusOK
		response.Status = "success"
	}
	return c.JSON(response.Code, response)
}
