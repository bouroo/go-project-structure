package repository

import (
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/entity"
	"gorm.io/gorm/clause"
)

func CreateUserAccount(user *entity.UserAccount) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	if err = dbTx.Omit(clause.Associations).Create(user).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func ReadUserAccount(userID string) (user entity.UserAccount, err error) {
	err = datasources.DBConn.Where(entity.UserAccount{ID: userID}).First(&user).Error
	return
}

func UpdateUserAccount(userID string, user entity.UserAccount) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Model(&entity.UserAccount{})
	dbTx.Where(entity.UserAccount{ID: userID})

	if err = dbTx.Updates(user).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func DeleteUserAccount(userID string) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Where(entity.UserAccount{ID: userID})
	if err = dbTx.Delete(&entity.UserAccount{}).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}
