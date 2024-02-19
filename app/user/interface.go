package user

import "github.com/bouroo/go-clean-arch/entity"

type Repository interface {
	CreateUserAccount(user *entity.UserAccount) (err error)
	ReadUserAccount(userID, username, email string) (user entity.UserAccount, err error)
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

	ReadUserDetails(userID, username, email string) (user entity.UserAccount, err error)
}

type Usecase interface {
	CreateUserAccount(user entity.UserAccount) (err error)
	ReadUserAccount(userID, username, email string) (user entity.UserAccount, err error)
	UpdateUserAccount(userID string, user entity.UserAccount) (err error)
	DeleteUserAccount(userID string) (err error)

	ReadUserDetails(userID, username, email string) (user entity.UserAccount, err error)
}

type Handler interface {
}
