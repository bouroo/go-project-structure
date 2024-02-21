package repository

import (
	"github.com/bouroo/go-clean-arch/entity"
)

func (r *userRepository) MigrateTable() (err error) {
	return r.db.AutoMigrate(
		&entity.UserAccount{},
		&entity.UserProfile{},
		&entity.UserAddress{},
	)
}
