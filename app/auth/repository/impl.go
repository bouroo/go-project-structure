package repository

import (
	"github.com/bouroo/go-clean-arch/app/auth"
	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{
		DB: db,
	}
}
