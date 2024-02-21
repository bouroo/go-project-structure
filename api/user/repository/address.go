package repository

import (
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/entity"
	"gorm.io/gorm/clause"
)

func CreateUserAddress(address *entity.UserAddress) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	if err = dbTx.Omit(clause.Associations).Create(address).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func ReadUserAddress(addressID, userID string) (addresses []entity.UserAddress, err error) {
	dbTx := datasources.DBConn.Model(&entity.UserAddress{})
	if len(addressID) != 0 {
		dbTx.Where(entity.UserAddress{ID: addressID})
	} else if len(userID) != 0 {
		dbTx.Where(entity.UserAddress{UserID: userID})
	}
	err = dbTx.Find(&addresses).Error
	return
}

func UpdateUserAddress(addressID string, address entity.UserAddress) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Model(&entity.UserAddress{})

	dbTx.Where(entity.UserAddress{ID: addressID})

	if err = dbTx.Updates(address).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func DeleteUserAddress(addressID string) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Where(entity.UserAddress{ID: addressID})
	if err = dbTx.Delete(&entity.UserAddress{}).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}
