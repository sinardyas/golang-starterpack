package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BookCode string `gorm:"type:binary(16);" json:"book_code"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Body     string `gorm:"type:text" json:"body"`
	UserID   int    `gorm:"index" json:"-"`
}
