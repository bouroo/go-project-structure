package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserAccount struct {
	ID        string         `gorm:"size:26;primary_key"`
	Email     string         `gorm:"size:50;uniqueIndex:udx_email"`
	Password  string         `gorm:"size:100"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:udx_email"`
	// Reference
	UserProfile UserProfile   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserAddress []UserAddress `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (UserAccount) TableName() string {
	return "tb_user_account"
}

func (u *UserAccount) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = ulid.Make().String()
	return
}

type UserProfile struct {
	ID        string         `gorm:"size:26;primary_key"`
	UserID    string         `gorm:"size:26;index"`
	FirstName string         `gorm:"size:100"`
	LastName  string         `gorm:"size:100"`
	Phone     string         `gorm:"size:20"`
	Avatar    string         `gorm:"size:255"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:udx_email"`
}

func (UserProfile) TableName() string {
	return "tb_user_profile"
}

func (u *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = ulid.Make().String()
	return
}

type UserAddress struct {
	ID        string         `gorm:"size:26;primary_key"`
	UserID    string         `gorm:"size:26;index"`
	Number    string         `gorm:"size:50"`
	Street    string         `gorm:"size:50"`
	City      string         `gorm:"size:50"`
	Province  string         `gorm:"size:50"`
	Country   string         `gorm:"size:30"`
	PostCode  string         `gorm:"size:16"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserAddress) TableName() string {
	return "tb_user_address"
}

func (u *UserAddress) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = ulid.Make().String()
	return
}
