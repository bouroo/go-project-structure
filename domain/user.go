package domain

import (
	"github.com/bouroo/go-clean-arch/entity"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	MigrateTable() (err error)

	CreateUserAccount(user *entity.UserAccount) (err error)
	ReadUserAccount(userID, email string) (user entity.UserAccount, err error)
	UpdateUserAccount(userID string, user entity.UserAccount) (err error)
	DeleteUserAccount(userID string) (err error)

	CreateUserProfile(userProfile *entity.UserProfile) (err error)
	ReadUserProfile(userID string) (userProfile entity.UserProfile, err error)
	UpdateUserProfile(userID string, userProfile entity.UserProfile) (err error)
	DeleteUserProfile(userID string) (err error)

	CreateUserAddress(userAddress *entity.UserAddress) (err error)
	ReadUserAddress(userID, addressID string) (addresses []entity.UserAddress, err error)
	UpdateUserAddress(userAddressID string, userAddress entity.UserAddress) (err error)
	DeleteUserAddress(userAddressID string) (err error)

	ReadUserDetails(userID, email string) (user entity.UserAccount, err error)
}

type UserUsecase interface {
	CreateUserAccount(user *entity.UserAccount) (err error)
	ReadUserAccount(userID, email string) (user entity.UserAccount, err error)
	UpdateUserAccount(userID string, user entity.UserAccount) (err error)
	DeleteUserAccount(userID string) (err error)

	CreateUserProfile(userProfile *entity.UserProfile) (err error)
	ReadUserProfile(userID string) (userProfile entity.UserProfile, err error)
	UpdateUserProfile(userID string, userProfile entity.UserProfile) (err error)
	DeleteUserProfile(userID string) (err error)

	CreateUserAddress(userAddress *entity.UserAddress) (err error)
	ReadUserAddress(userID, addressID string) (addresses []entity.UserAddress, err error)
	UpdateUserAddress(userAddressID string, userAddress entity.UserAddress) (err error)
	DeleteUserAddress(userAddressID string) (err error)

	ReadUserDetails(userID, email string) (user entity.UserAccount, err error)
}

type UserHandler interface {
	RegisterRoute(e *echo.Echo) *echo.Echo
}
