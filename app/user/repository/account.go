package repository

import (
	"github.com/bouroo/go-clean-arch/entity"
	"gorm.io/gorm/clause"
)

func (r *userRepository) CreateUserAccount(user *entity.UserAccount) (err error) {

	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Omit(clause.Associations).Create(user).Error
	if err != nil {
		r.Logger.Error("CreateUserAccount", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) ReadUserAccount(userID, username, email string) (user entity.UserAccount, err error) {

	selected := struct {
		Username string
		Password string
	}{}

	dbTx := r.DB.Model(&entity.UserAccount{})

	dbTx.Select(&selected)

	if len(userID) != 0 {
		dbTx.Where(entity.UserAccount{ID: userID})
	} else if len(email) != 0 {
		dbTx.Where(entity.UserAccount{Email: email})
	}

	if len(username) != 0 {
		dbTx.Where(entity.UserAccount{Email: username})
	}

	err = dbTx.First(&user).Error
	if err != nil {
		r.Logger.Error("ReadUserAccount", "error", err)
		return user, err
	}

	return user, err
}

func (r *userRepository) UpdateUserAccount(userID string, userAccount entity.UserAccount) (err error) {

	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Model(&entity.UserAccount{ID: userID}).Omit(clause.Associations).Updates(userAccount).Error
	if err != nil {
		r.Logger.Error("UpdateUserAccount", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) DeleteUserAccount(userID string) (err error) {
	dbTx := r.DB.Begin()
	defer dbTx.Rollback()

	err = dbTx.Delete(&entity.UserAccount{ID: userID}).Error
	if err != nil {
		r.Logger.Error("DeleteUserAccount", "error", err)
		return err
	}

	err = dbTx.Commit().Error
	return
}
