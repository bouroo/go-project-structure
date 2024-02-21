package repository

import (
	"github.com/bouroo/go-clean-arch/entity"
)

func (r *userRepository) CreateUserProfile(userProfile *entity.UserProfile) (err error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err = dbTx.Create(userProfile).Error
	if err != nil {
		r.logger.Error("CreateUserProfile", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) ReadUserProfile(userID string) (userProfile entity.UserProfile, err error) {
	err = r.db.First(&userProfile, userID).Error
	if err != nil {
		r.logger.Error("ReadUserProfile", "error", err)
		return userProfile, err
	}
	return
}

func (r *userRepository) UpdateUserProfile(userID string, userProfile entity.UserProfile) (err error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err = dbTx.Model(&entity.UserProfile{ID: userID}).Updates(userProfile).Error
	if err != nil {
		r.logger.Error("UpdateUserProfile", "error", err)
		return err
	}

	return dbTx.Commit().Error
}

func (r *userRepository) DeleteUserProfile(userID string) (err error) {
	dbTx := r.db.Begin()
	defer dbTx.Rollback()

	err = dbTx.Delete(&entity.UserProfile{}, userID).Error
	if err != nil {
		r.logger.Error("DeleteUserProfile", "error", err)
		return err
	}

	return dbTx.Commit().Error
}
