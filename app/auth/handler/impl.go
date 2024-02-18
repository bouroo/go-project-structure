package handler

import "github.com/bouroo/go-clean-arch/app/auth"

type authHandler struct {
	authUsecase auth.Usecase
}

func NewAuthHandler(authUsecase auth.Usecase) auth.Handler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}
