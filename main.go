package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect
	"github.com/sinardyas/golang-crud/controllers"
	"github.com/sinardyas/golang-crud/database"
)

func main() {
	router := gin.Default()
	user := router.Group("user")

	db := database.Init()

	ctl := new(controllers.User)
	ctl.Init(db, user)
	router.Run()
}
