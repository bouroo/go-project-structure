package repository

import "github.com/bouroo/go-clean-arch/internal/entity"

func (r *userRepository) ReadUserDetails(userID, email string) (user entity.UserAccount, err error) {

	dbTx := r.db.Model(&entity.UserAccount{})

	dbTx.Joins("UserProfile")

	dbTx.Joins("UserAddress")

	if len(userID) != 0 {
		dbTx.Where(entity.UserAccount{ID: userID})
	} else if len(email) != 0 {
		dbTx.Where(entity.UserAccount{Email: email})
	}

	err = dbTx.First(&user).Error
	if err != nil {
		r.logger.Error("ReadUserDetails", "error", err)
		return user, err
	}

	return user, err
}
