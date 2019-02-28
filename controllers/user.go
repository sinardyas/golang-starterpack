package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sinardyas/golang-crud/helper"
	"github.com/sinardyas/golang-crud/models"

	"github.com/jinzhu/gorm"
)

// UserController model
type UserController struct {
	db  *gorm.DB
	err error
}

// Init : constructor
func Init(gormDB *gorm.DB, r *mux.Router) {
	userController := UserController{
		db: gormDB,
	}

	r.HandleFunc("/", userController.Get).Methods("GET")
	r.HandleFunc("/", userController.Create).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", userController.Update).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", userController.Delete).Methods("DELETE")
}

// Get : return list of all user
func (userController UserController) Get(res http.ResponseWriter, req *http.Request) {
	result := userController.db.Find(&[]models.User{})

	response := &helper.Response{}
	response.ResponseHandling(res, 200, true, result)
}

// Create : create new user
func (userController UserController) Create(res http.ResponseWriter, req *http.Request) {
	user := models.User{
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
	}

	result := userController.db.Create(&user)
	response := &helper.Response{}
	response.ResponseHandling(res, 200, true, result)
}

// Update : update record
func (userController UserController) Update(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	userController.db.First(&userModel, userID)

	updateParam := models.User{
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
	}

	result := userController.db.Model(&userModel).Updates(updateParam)
	response := &helper.Response{}
	response.ResponseHandling(res, 200, true, result)
}

// Delete : delete record
func (userController UserController) Delete(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	userController.db.First(&userModel, userID)

	result := userController.db.Delete(&userModel)
	response := &helper.Response{}
	response.ResponseHandling(res, 200, true, result)
}
