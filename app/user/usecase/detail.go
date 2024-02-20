package usecase

import "github.com/bouroo/go-clean-arch/entity"

func (u *userUsecase) ReadUserDetails(userID, username, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserDetails(userID, username, email)
}
