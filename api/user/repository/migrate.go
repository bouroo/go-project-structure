package repository

import (
	"github.com/bouroo/go-project-structure/datasources"
	"github.com/bouroo/go-project-structure/pkg/entity"
)

func MigrateTable() (err error) {
	return datasources.DBConn.AutoMigrate(
		&entity.UserAccount{},
		&entity.UserProfile{},
		&entity.UserAddress{},
	)
}
