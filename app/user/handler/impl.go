package handler

import "github.com/bouroo/go-clean-arch/app/user"

type userHandler struct {
	userUsecase user.Usecase
}

func NewUserHandler(userUsecase user.Usecase) user.Handler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}
