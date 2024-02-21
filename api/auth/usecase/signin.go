package usecase

import "github.com/bouroo/go-clean-arch/internal/entity"

func (u *authUsecase) SignIn(email string, password string) (user *entity.UserAccount, err error) {
	u.userRepo.ReadUserAccount("", email)
	return
}
