package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/routers"
	"github.com/spf13/viper"
)

func main() {
	router := mux.NewRouter()
	config.ServiceConf()

	var route routers.Route
	var database config.Database
	errors := database.DatabaseInit()
	if errors == nil {
		fmt.Println("Error when connecting to DB")
	}

	route.UserRouterHandling(router)
	route.BookRouterHandling(router)

	fmt.Println("Server started at localhost:" + viper.GetString("PORT"))
	http.ListenAndServe("0.0.0.0:"+viper.GetString("PORT"), router)
}
