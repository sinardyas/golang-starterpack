package controllers

import (
	"net/http"

	"github.com/sinardyas/golang-crud/middlewares"

	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/helpers"

	"gopkg.in/go-playground/validator.v9"

	"github.com/gorilla/mux"

	"github.com/sinardyas/golang-crud/models"
)

var validate = validator.New()
var response helpers.Response
var validateRequest helpers.ValidationRequest
var passwordHandling helpers.Password
var auth middlewares.Auth
var db config.Database

// UserController model
type UserController struct{}

// Get : return list of all user
func (*UserController) Get(res http.ResponseWriter, req *http.Request) {
	result := db.DatabaseInit().Find(&[]models.User{})
	response.ResponseHandling(res, http.StatusOK, true, "Successfully get data", result)
}

// Create : create new user
func (*UserController) Create(res http.ResponseWriter, req *http.Request) {
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

	isExist := db.DatabaseInit().Where("user_name = ?", user.UserName).First(&models.User{})
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

	result := db.DatabaseInit().Create(&user)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully created", result)
}

// Update : update record
func (*UserController) Update(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	db.DatabaseInit().First(&userModel, userID)

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

	isExist := db.DatabaseInit().Where("user_name = ?", updateParam.UserName).First(&models.User{})
	if isExist.RowsAffected > 0 {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Username already exist", nil)
		return
	}

	result := db.DatabaseInit().Model(&userModel).Updates(updateParam)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully updated", result)
}

// Delete : delete record
func (*UserController) Delete(res http.ResponseWriter, req *http.Request) {
	var userModel models.User
	userID := mux.Vars(req)["id"]
	db.DatabaseInit().First(&userModel, userID)

	result := db.DatabaseInit().Delete(&userModel)
	response.ResponseHandling(res, http.StatusOK, true, "Successfully deleted", result)
}

// Login : login and get token auth
func (*UserController) Login(res http.ResponseWriter, req *http.Request) {
	var user models.User
	var book []models.Book
	loginParam := models.Login{
		UserName: req.FormValue("user_name"),
		Password: req.FormValue("password"),
	}

	isValid := validateRequest.ValidateHandling(loginParam)
	if isValid != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Validation error!", isValid)
		return
	}

	result := db.DatabaseInit().Where("user_name = ?", loginParam.UserName).Find(&user)
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

	errors := db.DatabaseInit().Model(&user).Related(&book).Error
	if errors != nil {
		response.ResponseHandling(res, http.StatusBadRequest, false, "Error occured!", errors)
		return
	}

	user.Books = book
	loginParam.Token = token
	loginParam.UserData = user
	response.ResponseHandling(res, http.StatusOK, true, "Logged in", loginParam)
}
