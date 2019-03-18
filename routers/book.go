package routers

import (
	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/controllers"
)

// TODO :
// separate router from controller

func (*Route) BookRouterHandling(router *mux.Router) {
	// var auth helpers.Auth
	var bookController controllers.BookController

	subrouter := router.PathPrefix("/book/").Subrouter()
	subrouter.HandleFunc("/", bookController.Get).Methods("GET")
}
