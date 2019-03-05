package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/gorilla/mux"

	"github.com/sinardyas/golang-crud/helper"
	"github.com/sinardyas/golang-crud/models"

	"github.com/jinzhu/gorm"
)

var validate = validator.New()
var response helper.Response
var validateRequest helper.ValidationRequest

// UserController model
type UserController struct {
	db  *gorm.DB
	err error
}

// Init : constructor
func Init(gormDB *gorm.DB, router *mux.Router) {
	userController := UserController{
		db: gormDB,
	}

	router.HandleFunc("/", userController.Get).Methods("GET")
	router.HandleFunc("/", userController.Create).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", userController.Update).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", userController.Delete).Methods("DELETE")
}

// Get : return list of all user
func (userController UserController) Get(res http.ResponseWriter, req *http.Request) {
	result := userController.db.Find(&[]models.User{})
	fmt.Println(result)

	response.ResponseHandling(res, http.StatusOK, true, result)
}

// Create : create new user
func (userController UserController) Create(res http.ResponseWriter, req *http.Request) {
	user := &models.User{
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
		UserName:  req.FormValue("userName"),
	}

	isValid := validateRequest.ValidateHandling(user)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, isValid)
		return
	}

	isExist := userController.db.Where("user_name = ?", user.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist!")
		return
	}

	result := userController.db.Create(&user)
	response.ResponseHandling(res, http.StatusOK, true, result)
}

// Update : update record
func (userController UserController) Update(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	userController.db.First(&userModel, userID)

	updateParam := models.User{
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
		UserName:  req.FormValue("userName"),
	}

	isValid := validateRequest.ValidateHandling(updateParam)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, isValid)
		return
	}

	isExist := userController.db.Where("user_name = ?", updateParam.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist!")
		return
	}

	result := userController.db.Model(&userModel).Updates(updateParam)
	response.ResponseHandling(res, http.StatusOK, true, result)
}

// Delete : delete record
func (userController UserController) Delete(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	userController.db.First(&userModel, userID)

	result := userController.db.Delete(&userModel)
	response.ResponseHandling(res, http.StatusOK, true, result)
}
