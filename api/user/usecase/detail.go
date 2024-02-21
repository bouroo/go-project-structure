package usecase

import "github.com/bouroo/go-clean-arch/internal/entity"

func (u *userUsecase) ReadUserDetails(userID, email string) (user entity.UserAccount, err error) {

	return u.userRepo.ReadUserDetails(userID, email)
}
