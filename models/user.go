package models

import (
	"github.com/jinzhu/gorm"
)

// User : presists user data in database
type User struct {
	gorm.Model
	FirstName       string `gorm:"size:255" validate:"required" json:"first_name"`
	LastName        string `gorm:"size:255" validate:"required" json:"last_name"`
	UserName        string `gorm:"size:30" validate:"required" json:"user_name"`
	Password        string `gorm:"size:255" validate:"required" json:"-"`
	ConfirmPassword string `gorm:"-" validate:"required" json:"-"`
}

type Login struct {
	UserName string `validate:"required" json:"user_name"`
	Password string `validate:"required" json:"-" json:"password"`
	Token    string
}
