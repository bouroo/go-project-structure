package usecase

import "github.com/bouroo/go-clean-arch/pkg/helper"

func (u *authUsecase) ResetPassword(email string) (token string, err error) {
	randBytes, err := helper.GenerateRandomBytes(8)
	if err != nil {
		return
	}
	token = helper.ComputeCRC64(randBytes)
	u.userRepo.ReadUserAccount("", email)
	return
}
