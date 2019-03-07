package routers

import (
	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/controllers"
	"github.com/sinardyas/golang-crud/helpers"
)

// TODO :
// separate router from controller
type UserRouter struct{}

func (*UserRouter) UserRouterHandling(router *mux.Router) {
	var auth helpers.Auth
	var userController controllers.UserController

	subrouter := router.PathPrefix("/user/").Subrouter()
	subrouter.HandleFunc("/", auth.MiddlewareAuth(userController.Get)).Methods("GET")
	subrouter.HandleFunc("/", auth.MiddlewareAuth(userController.Create)).Methods("POST")
	subrouter.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(userController.Update)).Methods("PUT")
	subrouter.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(userController.Delete)).Methods("DELETE")
	subrouter.HandleFunc("/login", userController.Login).Methods("POST")
}
