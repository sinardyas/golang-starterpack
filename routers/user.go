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

	router.HandleFunc("/", auth.MiddlewareAuth(userController.Get)).Methods("GET")
	router.HandleFunc("/", auth.MiddlewareAuth(userController.Create)).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(userController.Update)).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(userController.Delete)).Methods("DELETE")
	router.HandleFunc("/login", userController.Login).Methods("POST")
}
