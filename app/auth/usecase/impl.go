package usecase

import "github.com/bouroo/go-clean-arch/app/auth"

type authUsecase struct {
	authRepo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}
