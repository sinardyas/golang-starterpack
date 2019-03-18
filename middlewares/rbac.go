package middlewares

import (
	"net/http"

	"github.com/sinardyas/golang-crud/models"
)

type RBAC struct {
}

func (*RBAC) AccessLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authData := req.Context().Value("data").(models.User)

		if authData.Role != 1 {
			response.ResponseHandling(res, http.StatusUnauthorized, false, "Unauthorized", nil)
			return
		}

		next.ServeHTTP(res, req)
	})
}
