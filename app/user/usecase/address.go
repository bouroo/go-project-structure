package usecase

import "github.com/bouroo/go-clean-arch/entity"

func (u *userUsecase) CreateUserAddress(user *entity.UserAddress) (err error) {

	return u.userRepo.CreateUserAddress(user)
}

func (u *userUsecase) ReadUserAddress(userID, addressID string) (addresses []entity.UserAddress, err error) {

	return u.userRepo.ReadUserAddress(userID, addressID)
}

func (u *userUsecase) UpdateUserAddress(userID string, user entity.UserAddress) (err error) {

	return u.userRepo.UpdateUserAddress(userID, user)
}

func (u *userUsecase) DeleteUserAddress(userID string) (err error) {

	return u.userRepo.DeleteUserAddress(userID)
}
