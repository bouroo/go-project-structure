package repository

import (
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/entity"
)

func ReadUserDetails(userID, Email string) (user entity.UserAccount, err error) {
	dbTx := datasources.DBConn.Model(&entity.UserAccount{})
	dbTx.Joins("UserProfile")
	dbTx.Joins("UserAddress")

	if len(userID) != 0 {
		dbTx.Where(entity.UserAccount{ID: userID})
	} else if len(Email) != 0 {
		dbTx.Where(entity.UserAccount{Email: Email})
	}

	err = dbTx.First(&user).Error

	return
}
