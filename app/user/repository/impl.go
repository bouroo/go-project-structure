package repository

import (
	"github.com/bouroo/go-clean-arch/app/user"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		DB: db,
	}
}
