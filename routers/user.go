package routers

import (
	"github.com/gorilla/mux"
	"github.com/sinardyas/golang-crud/controllers"
	middlewares "github.com/sinardyas/golang-crud/middleware"
)

// TODO :
// separate router from controller
type Route struct{}

func (*Route) UserRouterHandling(router *mux.Router) {
	var auth middlewares.Auth
	var rbac middlewares.RBAC
	var userController controllers.UserController

	subrouter := router.PathPrefix("/user/").Subrouter()

	subrouter.Use(auth.MiddlewareAuth)
	subrouter.Use(rbac.AccessLimit)

	subrouter.HandleFunc("/", userController.Get).Methods("GET")
	subrouter.HandleFunc("/", userController.Create).Methods("POST")
	subrouter.HandleFunc("/{id:[0-9]+}", userController.Update).Methods("PUT")
	subrouter.HandleFunc("/{id:[0-9]+}", userController.Delete).Methods("DELETE")

	router.HandleFunc("/login", userController.Login).Methods("POST")
}
