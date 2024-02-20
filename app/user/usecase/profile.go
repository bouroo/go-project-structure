package usecase

import "github.com/bouroo/go-clean-arch/entity"

func (u *userUsecase) CreateUserProfile(user *entity.UserProfile) (err error) {

	return u.userRepo.CreateUserProfile(user)
}

func (u *userUsecase) ReadUserProfile(userID string) (user entity.UserProfile, err error) {

	return u.userRepo.ReadUserProfile(userID)
}

func (u *userUsecase) UpdateUserProfile(userID string, user entity.UserProfile) (err error) {

	return u.userRepo.UpdateUserProfile(userID, user)
}

func (u *userUsecase) DeleteUserProfile(userID string) (err error) {

	return u.userRepo.DeleteUserProfile(userID)
}
