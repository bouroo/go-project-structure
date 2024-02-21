package usecase

import (
	"github.com/bouroo/go-project-structure/api/user/repository"
	"github.com/bouroo/go-project-structure/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserAccount(user *entity.UserAccount) (err error) {
	return repository.CreateUserAccount(user)
}

func ReadUserAccount(userID string) (user entity.UserAccount, err error) {
	return repository.ReadUserAccount(userID)
}

func UpdateUserAccount(userID string, user entity.UserAccount) (err error) {
	return repository.UpdateUserAccount(userID, user)
}

func DeleteUserAccount(userID string) (err error) {
	return repository.DeleteUserAccount(userID)
}

func ChangePassword(userID string, password string) (err error) {

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	err = repository.UpdateUserAccount(userID, entity.UserAccount{Password: string(passwordByte)})
	return
}
