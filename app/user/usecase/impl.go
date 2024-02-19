package usecase

import (
	"log/slog"

	"github.com/bouroo/go-clean-arch/app/user"
	"github.com/bouroo/go-clean-arch/entity"
)

type userUsecase struct {
	userRepo user.Repository
	Logger   *slog.Logger
}

func NewUserUsecase(userRepo user.Repository, logger *slog.Logger) user.Usecase {
	return &userUsecase{
		userRepo: userRepo,
		Logger:   logger,
	}
}

func (u *userUsecase) CreateUserAccount(user entity.UserAccount) (err error) {

	return u.userRepo.CreateUserAccount(&user)
}

func (u *userUsecase) ReadUserAccount(userID, username, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserAccount(userID, username, email)
}

func (u *userUsecase) UpdateUserAccount(userID string, user entity.UserAccount) (err error) {

	return u.userRepo.UpdateUserAccount(userID, user)
}

func (u *userUsecase) DeleteUserAccount(userID string) (err error) {

	return u.userRepo.DeleteUserAccount(userID)
}

func (u *userUsecase) ReadUserDetails(userID, username, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserDetails(userID, username, email)
}
