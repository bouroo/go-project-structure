package repository

import (
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/entity"
	"gorm.io/gorm/clause"
)

func CreateUserProfile(profile *entity.UserProfile) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	if err = dbTx.Omit(clause.Associations).Create(profile).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func ReadUserProfile(profileID, userID string) (profile entity.UserProfile, err error) {
	dbTx := datasources.DBConn.Model(&entity.UserProfile{})
	if len(profileID) != 0 {
		dbTx.Where(entity.UserProfile{ID: profileID})
	} else if len(userID) != 0 {
		dbTx.Where(entity.UserProfile{UserID: userID})
	}
	err = dbTx.First(&profile).Error
	return
}

func UpdateUserProfile(profileID string, userID string, profile entity.UserProfile) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Model(&entity.UserProfile{})

	if len(profileID) != 0 {
		dbTx.Where(entity.UserProfile{ID: profileID})
	} else if len(userID) != 0 {
		dbTx.Where(entity.UserProfile{UserID: userID})
	}

	if err = dbTx.FirstOrCreate(&profile).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}

func DeleteUserProfile(profileID string) (err error) {
	dbTx := datasources.DBConn.Begin()
	defer dbTx.Commit()

	dbTx.Where(entity.UserProfile{ID: profileID})
	if err = dbTx.Delete(&entity.UserProfile{}).Error; err != nil {
		return
	}

	return dbTx.Commit().Error
}
