package usecase

import "github.com/bouroo/go-clean-arch/entity"

func (u *userUsecase) CreateUserAccount(user *entity.UserAccount) (err error) {

	return u.userRepo.CreateUserAccount(user)
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
