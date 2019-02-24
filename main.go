package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect
	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/controllers"
)

func main() {
	router := gin.Default()
	userRoute := router.Group("user")

	db := config.Init()

	controllers.Init(db, userRoute)
	router.Run()
}
