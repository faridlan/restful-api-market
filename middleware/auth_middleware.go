package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/faridlan/restful-api-market/app"
	"github.com/faridlan/restful-api-market/exception"
	"github.com/faridlan/restful-api-market/helper"
	"github.com/faridlan/restful-api-market/model/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authorizationHeader := request.Header.Get("Authorization")

	if request.URL.Path == "/api/login" || request.URL.Path == "/api/register" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	//cek db
	db := app.NewDB()
	SQL := "select id,token from blacklist where token = ?"
	rows, err := db.QueryContext(context.Background(), SQL, tokenString)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		//UNAUTHORIZED 401
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	} else {
		//OK 200
		var claim = &web.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != web.JwtSigningMEethod {
				return nil, fmt.Errorf("signing method invalid")
			}
			return web.JwtSecret, nil
		})

		if err != nil {
			panic(exception.NewUnauthError(err.Error()))
		}

		if !token.Valid {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
			}

			helper.WriteToResponseBody(writer, webResponse)
			writer.WriteHeader(http.StatusBadRequest)
			return
		} else {
			middleware.Handler.ServeHTTP(writer, request)
		}
	}

}
