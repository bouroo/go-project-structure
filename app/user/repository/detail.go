package repository

import "github.com/bouroo/go-clean-arch/entity"

func (r *userRepository) ReadUserDetails(userID, username, email string) (user entity.UserAccount, err error) {

	dbTx := r.DB.Model(&entity.UserAccount{})

	dbTx.Joins("UserProfile", r.DB.Where(&entity.UserProfile{Email: email}))

	dbTx.Joins("UserAddress")

	if len(userID) != 0 {
		dbTx.Where(entity.UserAccount{ID: userID})
	}

	if len(username) != 0 {
		dbTx.Where(entity.UserAccount{Username: username})
	}

	err = dbTx.First(&user).Error
	if err != nil {
		r.Logger.Error("ReadUserDetails", "error", err)
		return user, err
	}

	return user, err
}
