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
var passwordHandling helper.Password
var auth helper.Auth
var db *gorm.DB

// UserController model
type UserController struct{}

// Init : constructor
func Init(gormDB *gorm.DB, router *mux.Router) {
	fmt.Println("Ïnit Function")
	db = gormDB

	router.HandleFunc("/", auth.MiddlewareAuth(Get, db)).Methods("GET")
	router.HandleFunc("/", auth.MiddlewareAuth(Create, db)).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(Update, db)).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", auth.MiddlewareAuth(Delete, db)).Methods("DELETE")
	router.HandleFunc("/login", Login).Methods("POST")
}

// Get : return list of all user
func Get(res http.ResponseWriter, req *http.Request) {
	result := db.Find(&[]models.User{})

	response.ResponseHandling(res, http.StatusOK, true, "Successfully get data", result)
}

// Create : create new user
func Create(res http.ResponseWriter, req *http.Request) {
	user := &models.User{
		FirstName:       req.FormValue("first_name"),
		LastName:        req.FormValue("last_name"),
		UserName:        req.FormValue("user_name"),
		Password:        req.FormValue("password"),
		ConfirmPassword: req.FormValue("confirm_password"),
	}

	isValid := validateRequest.ValidateHandling(user)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	isExist := db.Where("user_name = ?", user.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist", nil)
		return
	}

	if user.Password != user.ConfirmPassword {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Password not match", nil)
		return
	}

	hashedPassword := passwordHandling.HashAndSalt([]byte(user.Password))
	user.Password = hashedPassword

	result := db.Create(&user)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully created", result)
}

// Update : update record
func Update(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	db.First(&userModel, userID)

	updateParam := models.User{
		FirstName: req.FormValue("first_name"),
		LastName:  req.FormValue("last_name"),
		UserName:  req.FormValue("user_name"),
	}

	isValid := validateRequest.ValidateHandling(updateParam)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	isExist := db.Where("user_name = ?", updateParam.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist", nil)
		return
	}

	result := db.Model(&userModel).Updates(updateParam)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully updated", result)
}

// Delete : delete record
func Delete(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	db.First(&userModel, userID)

	result := db.Delete(&userModel)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully deleted", result)
}

// Login : login and get token auth
func Login(res http.ResponseWriter, req *http.Request) {
	var user models.User
	loginParam := models.Login{
		UserName: req.FormValue("user_name"),
		Password: req.FormValue("password"),
	}

	isValid := validateRequest.ValidateHandling(loginParam)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	result := db.Where("user_name = ?", loginParam.UserName).Find(&user)
	if result.RowsAffected == 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "User not found", nil)
		return
	}

	isPasswordMatch := passwordHandling.ComparePassword(user.Password, []byte(loginParam.Password))
	if isPasswordMatch == false {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Wrong password", nil)
		return
	}

	token, err := auth.GenerateToken(&user)
	if err != nil {
		response.ResponseHandling(res, http.StatusInternalServerError, false, "Login failed", err)
		return
	}

	loginParam.Token = token
	response.ResponseHandling(res, http.StatusOK, true, "Logged in", loginParam)
}
