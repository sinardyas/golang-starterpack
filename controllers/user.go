package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect
)

var db *gorm.DB
var err error

type User struct {
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	gorm.Model
}

func (u User) Init(gormDB *gorm.DB, r *gin.RouterGroup) {
	db = gormDB

	r.GET("/", u.GetList)
	r.POST("/", u.RegisterUser)
}

func (u User) GetList(c *gin.Context) {
	data := db.Find(&[]User{})

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func (u User) RegisterUser(c *gin.Context) {
	user := User{
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
	}

	result := db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result.Value})
}
