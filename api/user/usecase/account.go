package usecase

import (
	"github.com/bouroo/go-clean-arch/internal/entity"
	"github.com/bouroo/go-clean-arch/pkg/helper"
)

func (u *userUsecase) CreateUserAccount(user *entity.UserAccount) (err error) {

	_, err = u.userRepo.ReadUserAccount("", user.Email)

	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	return u.userRepo.CreateUserAccount(user)
}

func (u *userUsecase) ReadUserAccount(userID, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserAccount(userID, email)
}

func (u *userUsecase) UpdateUserAccount(userID string, user entity.UserAccount) (err error) {

	return u.userRepo.UpdateUserAccount(userID, user)
}

func (u *userUsecase) DeleteUserAccount(userID string) (err error) {

	return u.userRepo.DeleteUserAccount(userID)
}
