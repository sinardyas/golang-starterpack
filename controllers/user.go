package controllers

import (
	"net/http"

	"github.com/sinardyas/golang-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserController model
type UserController struct {
	db  *gorm.DB
	err error
}

// Init : constructor
func Init(gormDB *gorm.DB, router *gin.RouterGroup) {
	userController := UserController{
		db: gormDB,
	}

	router.GET("/", userController.GetList)
	router.POST("/", userController.RegisterUser)
}

// GetList : return list of all user
func (userController UserController) GetList(context *gin.Context) {
	data := userController.db.Find(&[]models.User{})

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

// RegisterUser : create new user
func (userController UserController) RegisterUser(context *gin.Context) {
	user := models.User{
		FirstName: context.PostForm("firstName"),
		LastName:  context.PostForm("lastName"),
	}

	result := userController.db.Create(&user)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result.Value})
}
