package models

import (
	"github.com/jinzhu/gorm"
)

// User : presists user data in database
type User struct {
	gorm.Model
	FirstName string `gorm:"size:255" validate:"required"`
	LastName  string `gorm:"size:255" validate:"required"`
	UserName  string `gorm:"size:30" validate:"required"`
}
