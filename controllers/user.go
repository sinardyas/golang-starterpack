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
var db *gorm.DB

// UserController model
type UserController struct{}

// Init : constructor
func Init(gormDB *gorm.DB, router *mux.Router) {
	db = gormDB

	router.HandleFunc("/", Get).Methods("GET")
	router.HandleFunc("/", Create).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", Update).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", Delete).Methods("DELETE")
	router.HandleFunc("/login", Login).Methods("POST")
}

// Get : return list of all user
func Get(res http.ResponseWriter, req *http.Request) {
	result := db.Find(&[]models.User{})
	fmt.Println(result)

	response.ResponseHandling(res, http.StatusOK, true, "Successfully get data", result)
}

// Create : create new user
func Create(res http.ResponseWriter, req *http.Request) {
	user := &models.User{
		FirstName:       req.FormValue("firstName"),
		LastName:        req.FormValue("lastName"),
		UserName:        req.FormValue("userName"),
		Password:        req.FormValue("password"),
		ConfirmPassword: req.FormValue("confirmPassword"),
	}

	isValid := validateRequest.ValidateHandling(user)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	isExist := db.Where("user_name = ?", user.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist!", nil)
		return
	}

	if user.Password != user.ConfirmPassword {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Password not match!", nil)
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
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
		UserName:  req.FormValue("userName"),
	}

	isValid := validateRequest.ValidateHandling(updateParam)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	isExist := db.Where("user_name = ?", updateParam.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist!", nil)
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

func Login(res http.ResponseWriter, req *http.Request) {
	var user models.User
	loginParam := models.User{
		UserName: req.FormValue("userName"),
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

	response.ResponseHandling(res, http.StatusOK, true, "Logged in", result)
}
