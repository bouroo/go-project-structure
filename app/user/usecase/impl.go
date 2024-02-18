package usecase

import "github.com/bouroo/go-clean-arch/app/user"

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
