package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/models"
	"github.com/sinardyas/golang-crud/routers"
	"github.com/spf13/viper"
)

func main() {
	router := mux.NewRouter()
	config.ServiceConf()

	var route routers.UserRouter
	var database config.Database
	db := database.DatabaseInit()
	db.AutoMigrate(&models.User{})

	route.UserRouterHandling(router)
	// controllers.Init(db, router)
	fmt.Println("Server started at localhost:" + viper.GetString("PORT"))
	http.ListenAndServe("0.0.0.0:"+viper.GetString("PORT"), router)
}
