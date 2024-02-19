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

func (u *userUsecase) CreateUser(user entity.UserAccount) (err error) {

	return u.userRepo.CreateUserAccount(&user)
}

func (u *userUsecase) GetUserAccount(userID, username, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserAccount(userID, username, email)
}

func (u *userUsecase) GetUserDetails(userID, username, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserDetails(userID, username, email)
}

func (u *userUsecase) UpdateUser(userID string, user entity.UserAccount) (err error) {

	return u.userRepo.UpdateUserAccount(userID, user)
}

func (u *userUsecase) DeleteUser(userID string) (err error) {

	return u.userRepo.DeleteUserAccount(userID)
}
