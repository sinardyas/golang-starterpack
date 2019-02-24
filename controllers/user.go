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

	router.GET("/", userController.Get)
	router.POST("/", userController.Create)
	router.PUT("/:id", userController.Update)
	router.DELETE("/:id", userController.Delete)
}

// Get : return list of all user
func (userController UserController) Get(context *gin.Context) {
	data := userController.db.Find(&[]models.User{})

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

// Create : create new user
func (userController UserController) Create(context *gin.Context) {
	user := models.User{
		FirstName: context.PostForm("firstName"),
		LastName:  context.PostForm("lastName"),
	}

	result := userController.db.Create(&user)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result.Value})
}

// Update : update record
func (userController UserController) Update(context *gin.Context) {
	var userModel models.User
	userID := context.Param("id")
	userController.db.First(&userModel, userID)

	updateParam := models.User{
		FirstName: context.PostForm("firstName"),
		LastName:  context.PostForm("lastName"),
	}

	result := userController.db.Model(&userModel).Updates(updateParam)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

// Delete : delete record
func (userController UserController) Delete(context *gin.Context) {
	var userModel models.User
	userID := context.Param("id")
	userController.db.First(&userModel, userID)

	result := userController.db.Delete(&userModel)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}
