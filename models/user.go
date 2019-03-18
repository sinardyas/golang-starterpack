package models

import (
	"github.com/jinzhu/gorm"
)

// User : presists user data in database
type User struct {
	gorm.Model
	FirstName       string `gorm:"size:255" json:"first_name"`
	LastName        string `gorm:"size:255" json:"last_name"`
	UserName        string `gorm:"size:30" json:"user_name"`
	Password        string `gorm:"size:255" json:"-"`
	ConfirmPassword string `gorm:"-" json:"-"`
	Role            uint   `json:"role_id"`
	Books           []Book `json:"books"`
}

type Login struct {
	UserName string `gorm:"-" json:"user_name"`
	Password string `gorm:"-" json:"-" json:"password"`
	Token    string `gorm:"-" json:"token"`
	UserData User   `gorm:"-" json:"user_data"`
}
