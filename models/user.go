package models

import "github.com/jinzhu/gorm"

// User : presists user data in database
type User struct {
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	gorm.Model
}
