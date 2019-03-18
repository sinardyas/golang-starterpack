package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/sinardyas/golang-crud/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/sinardyas/golang-crud/config"
	"github.com/sinardyas/golang-crud/models"
	"github.com/spf13/viper"
)

var response helpers.Response

type Auth struct {
	jwt.StandardClaims
	Data interface{} `json:"data"`
}

func (*Auth) GenerateToken(data interface{}) (string, error) {
	claim := Auth{
		StandardClaims: jwt.StandardClaims{
			Issuer:    viper.GetString("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Data: data,
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(viper.GetString("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (*Auth) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil {
			response.ResponseHandling(res, http.StatusBadRequest, false, "Failed to Authenticate", err.Error())
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response.ResponseHandling(res, http.StatusBadRequest, false, "Failed to Authenticate", nil)
			return
		}

		data := claims["data"].(map[string]interface{})

		var user models.User
		var db config.Database
		result := db.DatabaseInit().Where("user_name = ?", data["user_name"]).Find(&user)

		if result.RowsAffected == 0 {
			response.ResponseHandling(res, http.StatusBadRequest, false, "User not found", nil)
			return
		}

		ctx := context.WithValue(context.Background(), "data", user)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}
