package controllers

import (
	"fmt"
	"net/http"

	"github.com/sinardyas/golang-crud/models"
)

type BookController struct{}

func (*BookController) Get(res http.ResponseWriter, req *http.Request) {
	result := db.DatabaseInit().Model(&models.User{}).Related(&models.Book{})

	fmt.Println(result)

	response.ResponseHandling(res, http.StatusOK, true, "Successfully get data", "Success")
}
