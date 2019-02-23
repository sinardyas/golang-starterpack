package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect
	"github.com/sinardyas/golang-crud/controllers"
)

func init() {
	db, err := gorm.Open("mysql", "root:root@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&controllers.User{})
}

func main() {
	fmt.Println("Migrate RUN")
}
