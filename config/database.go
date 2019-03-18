package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect
	"github.com/sinardyas/golang-crud/models"
	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	url      string
}

func (database *Database) DatabaseInit() *gorm.DB {
	database.Host = viper.GetString("DB_HOST")
	database.Port = viper.GetString("DB_PORT")
	database.User = viper.GetString("DB_USER")
	database.Password = viper.GetString("DB_PASS")
	database.Name = viper.GetString("DB_NAME")
	database.url = database.User + ":" + database.Password + "@(" + database.Host + ":" + database.Port + ")/" + database.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", database.url)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})

	return db
}
