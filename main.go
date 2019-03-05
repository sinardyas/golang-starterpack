package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/controllers"
	"github.com/spf13/viper"
)

var database config.Database

func main() {
	router := mux.NewRouter()

	config.ServiceConf()
	db := database.DatabaseInit()
	controllers.Init(db, router)

	fmt.Println("Server started at localhost:" + viper.GetString("PORT"))
	http.ListenAndServe("0.0.0.0:"+viper.GetString("PORT"), router)
}
